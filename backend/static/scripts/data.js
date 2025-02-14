// API Endpoints
const API_ENDPOINTS = {
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

// userData holds authenticated user's data
const userData = {
    data:(() => {
        try {
            const user = localStorage.getItem('userdata');
            return user ? JSON.parse(user) : null;
        } catch (error) {
            console.error("Error retrieving user data:", error);
            return null;
        }
    })()
};

// Constants
const CONSTANTS = {
    MIN_PASSWORD_LENGTH: 8,
    MAX_PASSWORD_LENGTH: 35,
};

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

const commentLikeState = {
    comments: {}
};

const postLikeState = {
    posts: {},
    comments: {},
  };

const SAMPLE_COMMENTS = {
    "01": [
        {
            id: 1,
            author: "walter otieno",
            content: "Great tutorial! The step-by-step approach really helped me understand the concepts.",
            timeAgo: "1h ago",
            likes: 5,
            replies: [
                {
                    id: 101,
                    author: "james ochieng",
                    content: "Totally agree! The examples were very clear.",
                    timeAgo: "45m ago",
                    likes: 2
                }
            ]
        },
        {
            id: 2,
            author: "martin shikuku",
            content: "Would love to see a follow-up tutorial on authentication with Go",
            timeAgo: "45m ago",
            likes: 3,
            replies: []
        }
    ],
    "02": [
        {
            id: 3,
            author: "thagruok owino",
            content: "Versioning is crucial for API design. Good point about maintaining backwards compatibility.",
            timeAgo: "2h ago",
            likes: 7,
            replies: []
        }
    ],
    "03": [
        {
            id: 4,
            author: "grace neema",
            content: "The section about lazy loading really improved my site's performance. Thanks!",
            timeAgo: "30m ago",
            likes: 4,
            replies: []
        }
    ]
};

export {
    API_ENDPOINTS, userData, SAMPLE_POSTS,
    CONSTANTS, commentLikeState, postLikeState,
    SAMPLE_COMMENTS
};
