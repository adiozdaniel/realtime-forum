document.addEventListener('DOMContentLoaded', () => {
    // Initialize Lucide icons
    lucide.createIcons();
    // State Management
    const state = {
        currentView: 'overview',
        darkMode: localStorage.getItem('darkMode') === 'true',
        profilePic: localStorage.getItem('profilePic') || 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=300&h=300&fit=crop',
        bio: localStorage.getItem('userBio') || 'Hi, I love coding and sharing knowledge with the community!',
        posts: [
            { id: 1, content: 'Just learned about React hooks!', comments: 5, likes: 12, timestamp: '2h ago' },
            { id: 2, content: 'Working on a new project using TypeScript', comments: 3, likes: 8, timestamp: '5h ago' },
            { id: 3, content: 'Check out my latest blog post about web performance', comments: 8, likes: 15, timestamp: '1d ago' }
        ],
        userComments: [
            { id: 1, postTitle: 'Introduction to GraphQL', content: 'Great explanation! This helped me understand the concepts better.', likes: 5, timestamp: '3h ago' },
            { id: 2, postTitle: 'Docker Best Practices', content: "I've been using similar patterns in my projects. Very effective!", likes: 3, timestamp: '1d ago' }
        ],
        activities: [
            { type: 'post', content: 'Created a new post', timestamp: '2h ago' },
            { type: 'comment', content: "Commented on 'Docker Best Practices'", timestamp: '1d ago' },
            { type: 'like', content: "Liked 'Introduction to GraphQL'", timestamp: '1d ago' }
        ]
    };
    // DOM Elements
    const elements = {
        backButton: document.getElementById('backButton'),
        profileImage: document.getElementById('profileImage'),
        imageUpload: document.getElementById('imageUpload'),
        bioText: document.getElementById('bioText'),
        editBioButton: document.getElementById('editBioButton'),
        darkModeToggle: document.getElementById('darkModeToggle'),
        sections: {
            overview: document.getElementById('overviewSection'),
            posts: document.getElementById('postsSection'),
            comments: document.getElementById('commentsSection'),
            settings: document.getElementById('settingsSection')
        },
        sidebarItems: document.querySelectorAll('.sidebar-item')
    };
    // Initialize the app
    function init() {
        updateTheme();
        updateStats();
        renderActivities();
        renderPosts();
        renderComments();
        setupEventListeners();
        elements.bioText.textContent = state.bio;
        elements.profileImage.src = state.profilePic;
        updateActiveSection();
    }
    // Event Listeners
    function setupEventListeners() {
        elements.backButton.addEventListener('click', function(event) {
            window.location.href = "http://localhost:4000/"
        });
        // Dark Mode Toggle
        elements.darkModeToggle.addEventListener('click', toggleDarkMode);
        // Image Upload
        elements.imageUpload.addEventListener('change', handleImageUpload);
        // Bio Edit
        elements.editBioButton.addEventListener('click', () => {
            const newBio = prompt('Edit your bio:', state.bio);
            if (newBio) {
                state.bio = newBio;
                localStorage.setItem('userBio', newBio);
                elements.bioText.textContent = newBio;
            }
        });
        // Sidebar Navigation
        elements.sidebarItems.forEach(item => {
            item.addEventListener('click', () => {
                const view = item.dataset.view;
                switchView(view);
            });
        });
    }
    // View Management
    function switchView(view) {
        state.currentView = view;
        updateActiveSection();
    }
    function updateActiveSection() {
        // Hide all sections
        Object.values(elements.sections).forEach(section => {
            section.classList.add('hidden');
        });
        // Show current section
        elements.sections[state.currentView].classList.remove('hidden');
        // Update sidebar active state
        elements.sidebarItems.forEach(item => {
            item.classList.toggle('active', item.dataset.view === state.currentView);
        });
    }
    // UI Updates
    function toggleDarkMode() {
        state.darkMode = !state.darkMode;
        localStorage.setItem('darkMode', state.darkMode);
        updateTheme();
    }
    function updateTheme() {
        document.body.setAttribute('data-theme', state.darkMode ? 'dark' : 'light');
        elements.darkModeToggle.innerHTML = `<i data-lucide="${state.darkMode ? 'sun' : 'moon'}"></i>`;
        lucide.createIcons();
    }
    function handleImageUpload(e) {
        const file = e.target.files[0];
        if (file) {
            const reader = new FileReader();
            reader.onloadend = () => {
                state.profilePic = reader.result;
                localStorage.setItem('profilePic', reader.result);
                elements.profileImage.src = reader.result;
            };
            reader.readAsDataURL(file);
        }
    }
    function updateStats() {
        document.getElementById('postsCount').textContent = state.posts.length;
        document.getElementById('commentsCount').textContent = state.userComments.length;
        document.getElementById('likesCount').textContent =
            state.posts.reduce((acc, post) => acc + post.likes, 0);
    }
    function renderActivities() {
        const activityList = document.getElementById('activityList');
        activityList.innerHTML = state.activities.map(activity => `
            <div class="activity-item">
                <i data-lucide="clock"></i>
                <span>${activity.content}</span>
                <span>${activity.timestamp}</span>
            </div>
        `).join('');
        lucide.createIcons();
    }
    function renderPosts() {
        const postsList = document.getElementById('postsList');
        postsList.innerHTML = state.posts.map(post => `
            <div class="post-item">
                <p>${post.content}</p>
                <div class="post-actions">
                    <span><i data-lucide="thumbs-up"></i> ${post.likes}</span>
                    <span><i data-lucide="message-square"></i> ${post.comments}</span>
                    <span>${post.timestamp}</span>
                </div>
                <button class="delete-button" onclick="deletePost(${post.id})">
                    <i data-lucide="trash-2"></i>
                </button>
            </div>
        `).join('');
        lucide.createIcons();
    }
    function renderComments() {
        const commentsList = document.getElementById('commentsList');
        commentsList.innerHTML = state.userComments.map(comment => `
            <div class="comment-item">
                <h3>Re: ${comment.postTitle}</h3>
                <p>${comment.content}</p>
                <div class="post-actions">
                    <span><i data-lucide="thumbs-up"></i> ${comment.likes}</span>
                    <span>${comment.timestamp}</span>
                </div>
            </div>
        `).join('');
        lucide.createIcons();
    }
    // Initialize the app
    init();
    // Expose necessary functions to window for onclick handlers
    window.deletePost = (postId) => {
        state.posts = state.posts.filter(post => post.id !== postId);
        renderPosts();
        updateStats();
    };
});