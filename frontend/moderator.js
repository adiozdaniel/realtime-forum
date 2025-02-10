const initialData = {
    userReports: [
        { 
            id: 1, 
            reportedUser: 'mark', 
            reportedBy: 'joe',
            content: 'Inappropriate content posted in general discussion',
            postContent: 'This is the reported post content...',
            timeAgo: '5 minutes ago',
            category: 'General',
            reason: 'obscene',
            status: 'pending'
        },
        { 
            id: 2, 
            reportedUser: 'walter', 
            reportedBy: 'faith',
            content: 'Multiple spam messages in trading section',
            postContent: 'Buy now! Limited time offer...',
            timeAgo: '10 minutes ago',
            category: 'Trading',
            reason: 'spam',
            status: 'pending'
        },
        {
            id: 3,
            reportedUser: 'maya',
            reportedBy: 'violet',
            content: 'Harassing comments towards other users',
            postContent: 'Series of hostile comments...',
            timeAgo: '15 minutes ago',
            category: 'Discussion',
            reason: 'harassment',
            status: 'pending'
        }
    ],
    moderatorReports: [
        { 
            id: 1,
            reportedUser: 'victor',
            content: 'Multiple instances of harassment and threatening behavior',
            reportDate: '2024-02-10 14:30',
            timeAgo: '1 hour ago',
            status: 'sent',
            adminResponse: 'Under review - Admin Team'
        },
        {
            id: 2,
            reportedUser: 'otieno',
            content: 'Organized spam campaign across multiple sections',
            reportDate: '2024-02-10 13:15',
            timeAgo: '2 hours ago',
            status: 'resolved',
            adminResponse: 'User account suspended - Admin Team'
        }
    ]
};

function renderUserReports() {
    const reportsContainer = document.getElementById('user-reports-container');
    reportsContainer.innerHTML = initialData.userReports
        .filter(report => report.status === 'pending')
        .map(report => `
            <div class="report-card">
                <div class="report-info">
                    <div class="report-header">
                        <div class="report-header-main">
                            <span class="reported-user">Reported User: <strong>${report.reportedUser}</strong></span>
                            <span class="reporter">Reported By: ${report.reportedBy}</span>
                        </div>
                        <span class="time">${report.timeAgo}</span>
                    </div>
                    <div class="report-content">
                        <p class="report-description"><strong>Report Description:</strong> ${report.content}</p>
                        <div class="reported-content">
                            <strong>Reported Content:</strong>
                            <p class="content-preview">${report.postContent}</p>
                        </div>
                    </div>
                    <div class="report-meta">
                        <span class="category">Category: ${report.category}</span>
                        <span class="reason">Reason: ${report.reason}</span>
                    </div>
                </div>
                <div class="action-buttons">
                    <button class="report-btn" onclick="escalateToAdmin(${report.id})">
                        <i data-lucide="alert-triangle"></i>
                        Report to Admin
                    </button>
                    <button class="dismiss-btn" onclick="dismissReport(${report.id})">
                        <i data-lucide="x"></i>
                        Dismiss Report
                    </button>
                </div>
            </div>
    `).join('');
    lucide.createIcons();
}

function renderModeratorReports() {
    const reportsContainer = document.getElementById('moderator-reports-container');
    reportsContainer.innerHTML = initialData.moderatorReports.map(report => `
        <div class="report-card ${report.status === 'resolved' ? 'resolved' : ''}">
            <div class="report-info">
                <div class="report-header">
                    <div class="report-header-main">
                        <span class="reported-user">Reported User: <strong>${report.reportedUser}</strong></span>
                        <span class="status ${report.status}">${report.status.toUpperCase()}</span>
                    </div>
                    <span class="time">${report.timeAgo}</span>
                </div>
                <div class="report-content">
                    <p class="report-description">${report.content}</p>
                </div>
                <div class="report-status">
                    <span class="report-date">Reported on: ${report.reportDate}</span>
                    ${report.adminResponse ? `
                        <div class="admin-response">
                            <i data-lucide="message-square"></i>
                            Admin Response: ${report.adminResponse}
                        </div>
                    ` : ''}
                </div>
            </div>
        </div>
    `).join('');
    lucide.createIcons();
}

function escalateToAdmin(reportId) {
    const report = initialData.userReports.find(r => r.id === reportId);
    if (report) {
        report.status = 'escalated';
        const now = new Date();
        const formattedDate = now.toISOString().split('T')[0] + ' ' + 
                            now.toTimeString().split(' ')[0].substring(0, 5);
        
        initialData.moderatorReports.unshift({
            id: initialData.moderatorReports.length + 1,
            reportedUser: report.reportedUser,
            content: report.content,
            reportDate: formattedDate,
            timeAgo: 'Just now',
            status: 'sent'
        });
        renderUserReports();
        renderModeratorReports();
        showNotification('Report has been escalated to admin');
    }
}

function dismissReport(reportId) {
    const reportIndex = initialData.userReports.findIndex(r => r.id === reportId);
    if (reportIndex > -1) {
        initialData.userReports[reportIndex].status = 'dismissed';
        renderUserReports();
        showNotification('Report has been dismissed');
    }
}

function showNotification(message) {
    const notification = document.createElement('div');
    notification.className = 'notification';
    notification.innerHTML = `
        <i data-lucide="check-circle"></i>
        <span>${message}</span>
    `;
    document.body.appendChild(notification);
    lucide.createIcons();
    setTimeout(() => notification.remove(), 3000);
}
// Toggle dark mode
function toggleDarkMode() {
    document.body.classList.toggle('dark-mode');
    localStorage.setItem('darkMode', document.body.classList.contains('dark-mode'));
}

function switchSection(sectionId) {
    document.querySelectorAll('.section').forEach(section => {
        section.classList.add('hidden');
    });
    document.getElementById(`${sectionId}-section`).classList.remove('hidden');
    
    // Update active state of navigation buttons
    document.querySelectorAll('.category-button').forEach(button => {
        button.classList.remove('active');
        if (button.getAttribute('onclick').includes(sectionId)) {
            button.classList.add('active');
        }
    });
}

document.addEventListener('DOMContentLoaded', () => {
    renderUserReports();
    renderModeratorReports();
});