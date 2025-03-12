// Initial Data
const initialData = {
    reports: [
        { id: 1, user: 'shayo_victor', content: 'This is spam', timeAgo: '7 minutes ago' },
        { id: 2, user: 'martin_shikuku', content: 'This is spam', timeAgo: '2 days ago' }
    ],
    users: [
        { id: 1, username: 'jane_doe', joined: '2 years ago', lastActive: '3 days ago', posts: 40, likes: 100, isModerator: true },
        { id: 2, username: 'jim_smith', joined: '1 year ago', lastActive: '1 day ago', posts: 30, likes: 80, isModerator: true },
        { id: 3, username: 'john_doe', joined: '6 months ago', lastActive: '2 days ago', posts: 20, likes: 50, isModerator: false },
        { id: 4, username: 'jane_smith', joined: '3 months ago', lastActive: '1 week ago', posts: 10, likes: 20, isModerator: false }
    ],
    categories: [
        { id: 1, name: 'General', posts: 10, comments: 12 },
        { id: 2, name: 'Music', posts: 10, comments: 12 },
        { id: 3, name: 'Technology', posts: 10, comments: 12 },
        { id: 4, name: 'Movies', posts: 10, comments: 12 }
    ]
};

// Utility Functions
function getInitials(username) {
    return username
        .split('_')
        .map(word => word[0])
        .join('')
        .toUpperCase();
}

function getAvatarColor(username) {
    const colors = ['#3b82f6', '#10b981', '#8b5cf6', '#ec4899'];
    return colors[username.length % colors.length];
}

// DOM Elements
const reportsContainer = document.getElementById('reports-container');
const userList = document.getElementById('user-list');
const categoriesContainer = document.getElementById('categories-container');
const newCategoryInput = document.getElementById('new-category-input');
const addCategoryBtn = document.getElementById('add-category-btn');
const darkModeToggle = document.getElementById('darkModeToggle');

// Render Functions
function renderReports() {
    reportsContainer.innerHTML = initialData.reports.map(report => `
        <div class="post-card">
            <div class="user-info">
                <div class="user-avatar" style="background-color: ${getAvatarColor(report.user)}">
                    ${getInitials(report.user)}
                </div>
                <span>${report.user}</span>
            </div>
            <div class="post-meta">
                <span>${report.content}</span>
                <span>${report.timeAgo}</span>
            </div>
            <button class="actions-menu" onclick="toggleDropdown(event, ${report.id}, 'report')">
                <i data-lucide="more-horizontal"></i>
            </button>
            <div class="dropdown-menu hidden" id="dropdown-report-${report.id}">
                <!-- Add report-specific actions here -->
            </div>
        </div>
    `).join('');
    lucide.createIcons();
}

function renderUsers(filteredUsers = initialData.users) {
    userList.innerHTML = filteredUsers.map(user => `
        <tr>
            <td>
                <div class="user-info">
                    <div class="user-avatar" style="background-color: ${getAvatarColor(user.username)}">
                        ${getInitials(user.username)}
                    </div>
                    <span>${user.username}</span>
                </div>
            </td>
            <td>${user.joined}</td>
            <td>${user.lastActive}</td>
            <td>${user.posts}</td>
            <td>${user.likes}</td>
            <td>
                <button class="actions-menu" onclick="toggleDropdown(event, ${user.id}, 'user')">
                    <i data-lucide="more-horizontal"></i>
                </button>
                <div class="dropdown-menu hidden" id="dropdown-user-${user.id}">
                    <button onclick="handlePromoteDemote(${user.id})">${user.isModerator ? 'Demote' : 'Promote'} to Moderator</button>
                </div>
            </td>
        </tr>
    `).join('');
    lucide.createIcons();
}

function renderCategories() {
    categoriesContainer.innerHTML = initialData.categories.map(category => `
        <div class="post-card">
            <div class="category-content">
                <div>
                    <h3>${category.name}</h3>
                    <p class="post-meta">${category.posts} posts, ${category.comments} comments</p>
                </div>
                <button class="actions-menu" onclick="toggleDropdown(event, ${category.id}, 'category')">
                    <i data-lucide="more-horizontal"></i>
                </button>
                <div class="dropdown-menu hidden" id="dropdown-category-${category.id}">
                    <!-- Add category-specific actions here -->
                </div>
            </div>
        </div>
    `).join('');
    lucide.createIcons();
}

// Search Functionality
function searchUsers() {
    const searchInput = document.getElementById('user-search-input').value.toLowerCase();
    const filteredUsers = initialData.users.filter(user => 
        user.username.toLowerCase().includes(searchInput)
    );
    renderUsers(filteredUsers);
}

// Event Listener for Search Input
document.getElementById('user-search-input').addEventListener('input', searchUsers);

function handlePromoteDemote(id) {
    const user = initialData.users.find(user => user.id === id);
    if (user) {
        if (user.isModerator) {
            // Logic to demote the moderator
            user.isModerator = false;
            console.log(`Demoted moderator: ${user.username}`);
        } else {
            // Logic to promote the user to a moderator
            user.isModerator = true;
            console.log(`Promoted user to moderator: ${user.username}`);
        }
        renderUsers();
    }
}

function toggleDropdown(event, id, type) {
    event.stopPropagation();
    const dropdown = document.getElementById(`dropdown-${type}-${id}`);
    const allDropdowns = document.querySelectorAll('.dropdown-menu');
    allDropdowns.forEach(menu => {
        if (menu !== dropdown) {
            menu.classList.add('hidden');
        }
    });
    dropdown.classList.toggle('hidden');
}

function handleCategoryAction(categoryId) {
    console.log(`Action for category: ${categoryId}`);
    // Implement category edit/delete logic here
}

function handleAddCategory() {
    const categoryName = newCategoryInput.value.trim();
    if (categoryName) {
        initialData.categories.push({
            id: initialData.categories.length + 1,
            name: categoryName,
            posts: 0,
            comments: 0
        });
        newCategoryInput.value = '';
        renderCategories();
    }
}

function toggleDarkMode() {
    document.body.classList.toggle('dark-mode');
    const icon = darkModeToggle.querySelector('i');
    if (document.body.classList.contains('dark-mode')) {
        icon.setAttribute('data-lucide', 'sun');
    } else {
        icon.setAttribute('data-lucide', 'moon');
    }
    lucide.createIcons();
}

function switchSection(sectionId) {
    document.querySelectorAll('.section').forEach(section => {
        section.classList.add('hidden');
    });
    document.getElementById(`${sectionId}-section`).classList.remove('hidden');
}

// Event Listeners
addCategoryBtn.addEventListener('click', handleAddCategory);
darkModeToggle.addEventListener('click', toggleDarkMode);
document.addEventListener('click', () => {
    document.querySelectorAll('.dropdown-menu').forEach(menu => menu.classList.add('hidden'));
});

// Initialize the dashboard
function initDashboard() {
    renderReports();
    renderUsers();
    renderCategories();
    lucide.createIcons();
}

// Start the application
document.addEventListener('DOMContentLoaded', initDashboard);


const recentDiscussions = [
    {
        id: 1,
        title: "New Feature Request: Dark Mode for Code Editor",
        author: "sarah_dev",
        content: "I think it would be great if we could add dark mode support to the code editor...",
        timeAgo: "2 hours ago",
        replies: 12,
        likes: 24
    },
    {
        id: 2,
        title: "Bug Report: Mobile Navigation Issues",
        author: "tech_lead",
        content: "Users are reporting problems with the mobile navigation menu not responding...",
        timeAgo: "4 hours ago",
        replies: 8,
        likes: 15
    },
    {
        id: 3,
        title: "Documentation Update Needed",
        author: "doc_writer",
        content: "The API documentation needs to be updated to reflect recent changes...",
        timeAgo: "6 hours ago",
        replies: 5,
        likes: 10
    }
];


function renderRecentDiscussions() {
    const container = document.getElementById('recentDiscussionsContainer');
    if (!container) return;

    container.innerHTML = recentDiscussions.map(discussion => `
        <div class="discussion-card">
            <div class="discussion-header">
                <div>
                    <div class="user-info">
                        <div class="user-avatar" style="background-color: ${getAvatarColor(discussion.author)}">
                            ${getInitials(discussion.author)}
                        </div>
                        <span>${discussion.author}</span>
                    </div>
                    <h3 class="discussion-title">${discussion.title}</h3>
                </div>
                <button class="actions-menu" onclick="toggleDropdown(event, ${discussion.id}, 'discussion')">
                    <i data-lucide="more-horizontal"></i>
                </button>
            </div>
            <div class="discussion-content">${discussion.content}</div>
            <div class="discussion-footer">
                <div class="post-actions">
                    <button class="post-action-button">
                        <i data-lucide="message-square"></i>
                        <span>${discussion.replies}</span>
                    </button>
                    <button class="post-action-button">
                        <i data-lucide="heart"></i>
                        <span>${discussion.likes}</span>
                    </button>
                </div>
                <div class="discussion-meta">
                    <span>${discussion.timeAgo}</span>
                </div>
            </div>
        </div>
    `).join('');
    
    lucide.createIcons();
}

function initDashboard() {
    renderReports();
    renderUsers();
    renderCategories();
    renderRecentDiscussions(); 
    lucide.createIcons();
}