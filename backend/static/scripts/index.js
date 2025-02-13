import { PostService } from "./postsservice.js";
import { formatTimeAgo } from "./timestamps.js";

// API Endpoints
window.API_ENDPOINTS = {
  login: "/api/auth/login",
  register: "/api/auth/register",
  logout: "/api/auth/logout",
  check: "/api/auth/check",

  // posts ENDPOINTS
  allposts: "/api/posts",
  createpost: "/api/posts/create",
  deletepost: "/api/posts/delete",
  updatepost: "/api/posts/update",
  likepost: "/api/posts/like",
  dislikepost: "/api/posts/dislike",

  // comments ENDPOINTS
  listcommbypost: "/api/comments",
  createcomment: "/api/comments/create",
  deletecomment: "/api/comments/delete",
  updatecomment: "/api/comments/update",
  likecomment: "/api/comments/like",
  dislikecomment: "/api/comments/dislike",
};

// ResData object
window.RESDATA = {
  userData: (() => {
    try {
      const data = localStorage.getItem("res");
      return data ? JSON.parse(data) : null;
    } catch (error) {
      console.error("Error parsing localStorage data:", error);
      return null;
    }
  })(),
  profileImageElement: null,
};

// Constants
window.CONSTANTS = {
  MIN_PASSWORD_LENGTH: 8,
  MAX_PASSWORD_LENGTH: 35,
};

// DOM Elements
const postsContainer = document.querySelector("#postsContainer");

// Sample Data
const SAMPLE_POSTS = [
  {
    post_id: "01",
    post_title: "Getting Started with Go and Angular",
    post_author: "Jane Cooper",
    post_category: "Tutorial",
    post_likes: 42,
    post_comments: 12,
    post_content:
      "Learn how to build a modern web application using Go for the backend and Angular for the frontend...",
    post_timeAgo: "2h ago",
    post_hasComments: true,
  },
  {
    post_id: "02",
    post_title: "Best Practices for API Design",
    post_author: "John Smith",
    post_category: "Discussion",
    post_likes: 28,
    post_comments: 8,
    post_content:
      "Let's discuss the best practices for designing RESTful APIs that are both scalable and maintainable...",
    post_timeAgo: "4h ago",
    post_hasComments: true,
  },
  {
    post_id: "03",
    post_title: "Web Performance Optimization Tips",
    post_author: "Alice Johnson",
    post_category: "Guide",
    post_likes: 35,
    post_comments: 15,
    post_content:
      "Essential tips and tricks for optimizing your web application's performance...",
    post_timeAgo: "6h ago",
    post_hasComments: true,
  },
];

// Import comment functions
import { commentManager } from "./comment.js";

// PostManager class encapsulates post and comment management
class PostManager {
  constructor() {
    this.likeState = {
      posts: {},
      comments: {},
    };
    this.postService = new PostService();
  }

  // Create post HTML
  createPostHTML(post) {
    const isLiked = this.likeState.posts[post.post_id]?.likedBy.has("current-user");
    return `
      <article class="post-card" data-post-id="${post.post_id}">
        <div class="flex items-start justify-between">
          <div>
            <span class="post-category">${post.post_category}</span>
            <h3 class="post-title">${post.post_title}</h3>
            <p class="post-excerpt">${post.post_content}</p>
          </div>
        </div>
        <div class="post-footer">
          <div class="post-actions">
            <button class="post-action-button like-button ${isLiked ? "liked text-blue-600" : ""}" data-post-id="${post.post_id}">
              <i data-lucide="thumbs-up"></i>
              <span class="likes-count">${this.likeState.posts[post.post_id]?.count || 0}</span>
            </button>
            <button class="post-action-button comment-toggle" data-post-id="${post.post_id}">
              <i data-lucide="message-square"></i>
              <span class="comments-count">${post.post_comments}</span>
            </button>
          </div>
          <div class="post-meta">
            <span>by ${post.post_author}</span>
            <span>•</span>
            <span>${post.post_timeAgo}</span>
          </div>
        </div>
        <div class="comments-section hidden" id="comments-${post.post_id}">
          <div class="comments-container"></div>
          <form class="comment-form" data-post-id="${post.post_id}">
            <textarea placeholder="Write your comment..." class="comment-input"></textarea>
            <button type="submit" class="comment-submit">Post Comment</button>
          </form>
        </div>
      </article>
    `;
  }

  // Toggle comments section visibility
  toggleComments(e) {
    const commentButton = e.target.closest(".comment-toggle");
    if (!commentButton) return;

    const postId = commentButton.dataset.postId;
    const commentsSection = document.querySelector(`#comments-${postId}`);

    if (commentsSection.classList.contains("hidden")) {
      loadComments(postId);
    }

    commentsSection.classList.toggle("hidden");
  }

  // Render all posts
  renderPosts(posts = SAMPLE_POSTS) {
    postsContainer.innerHTML = posts.map((post) => this.createPostHTML(post)).join("");
    lucide.createIcons();

    this.attachPostEventListeners();
  }

  // Attach event listeners to post buttons
  attachPostEventListeners() {
    document.querySelectorAll(".like-button").forEach((button) => {
      button.addEventListener("click", this.handleLike.bind(this));
    });
    document.querySelectorAll(".comment-toggle").forEach((button) => {
      button.addEventListener("click", this.toggleComments.bind(this));
    });
    document.querySelectorAll(".comment-form").forEach((form) => {
      form.addEventListener("submit", commentManager.handleCommentSubmit);
    });
  }

  // Handle like button click for both posts and comments
  async handleLike(e) {
    const button = e.currentTarget.closest(".like-button");
    if (!button) return;
    const isComment = button.hasAttribute("data-comment-id");
    const postId = button.getAttribute("data-post-id");
    if (!postId) return;
    const commentId = isComment ? button.getAttribute("data-comment-id") : null;

    const stateRef = isComment ? (this.likeState.comments[postId] = this.likeState.comments[postId] || {}) : this.likeState.posts[postId];

    if (isComment && !stateRef[commentId]) {
      stateRef[commentId] = {
        count: parseInt(button.querySelector(".likes-count")?.textContent) || 0,
        likedBy: new Set(),
      };
    }
    const likeData = isComment ? stateRef[commentId] : stateRef;
    if (!likeData) return;

    if (!window.RESDATA?.userData) {
      window.location.href = "/auth";
      return;
    }

    const currentUserId = window.RESDATA.userData.user_id;
    const postData = { post_id: postId, user_id: currentUserId };

    if (likeData.likedBy.has(currentUserId)) {
      const res = await this.postService.dislikePost(postData);
      if (res.error) {
        alert(res.message);
        return;
      }
      likeData.count--;
      likeData.likedBy.delete(currentUserId);
      button.classList.remove("liked", "text-blue-600");
    } else {
      const res = await this.postService.likePost(postData);
      if (res.error) {
        alert(res.message);
        return;
      }
      likeData.count++;
      likeData.likedBy.add(currentUserId);
      button.classList.add("liked", "text-blue-600");
    }

    const likesCount = button.querySelector(".likes-count");
    if (likesCount) {
      likesCount.textContent = likeData.count;
    }

    button.classList.add("like-animation");
    setTimeout(() => button.classList.remove("like-animation"), 300);
  }

  // Initialize the application
  async init() {
    const posts = await this.postService.fetchPosts();
    const postList = Array.isArray(posts) ? posts : posts.data;

    postList.forEach((post) => SAMPLE_POSTS.push(post));

    fetchComments();

    SAMPLE_POSTS.forEach((post) => {
      post.post_timeAgo = formatTimeAgo(post.created_at);

      this.likeState.posts[post.post_id] = {
        count: post.post_likes,
        likedBy: new Set(),
      };
    });

    if (postsContainer) this.renderPosts();
  }
}

// Initialize
document.addEventListener("DOMContentLoaded", () => {
  const postManager = new PostManager();
  postManager.init();
});

export { PostManager, SAMPLE_POSTS };
