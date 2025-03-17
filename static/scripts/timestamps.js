export function formatTimeAgo(date) {
    const now = new Date();
    const timestamp = new Date(date);
    const secondsAgo = Math.floor((now - timestamp) / 1000);

    if (secondsAgo < 60) {
        return 'just now';
    }

    const minutesAgo = Math.floor(secondsAgo / 60);
    if (minutesAgo < 60) {
        return `${minutesAgo} ${minutesAgo === 1 ? 'minute' : 'minutes'} ago`;
    }

    const hoursAgo = Math.floor(minutesAgo / 60);
    if (hoursAgo < 24) {
        return `${hoursAgo} ${hoursAgo === 1 ? 'hour' : 'hours'} ago`;
    }

    const daysAgo = Math.floor(hoursAgo / 24);
    if (daysAgo < 7) {
        return `${daysAgo} ${daysAgo === 1 ? 'day' : 'days'} ago`;
    }

    const weeksAgo = Math.floor(daysAgo / 7);
    if (weeksAgo < 4) {
        return `${weeksAgo} ${weeksAgo === 1 ? 'week' : 'weeks'} ago`;
    }

    const monthsAgo = Math.floor(daysAgo / 30);
    if (monthsAgo < 12) {
        return `${monthsAgo} ${monthsAgo === 1 ? 'month' : 'months'} ago`;
    }

    const yearsAgo = Math.floor(daysAgo / 365);
    return `${yearsAgo} ${yearsAgo === 1 ? 'year' : 'years'} ago`;
}

// Update all timestamps on the page
export function updateTimestamps() {
    const timestampElements = document.querySelectorAll('[data-timestamp]');
    timestampElements.forEach(element => {
        const timestamp = element.getAttribute('data-timestamp');
        element.textContent = formatTimeAgo(timestamp);
    });
}

// Update timestamps every minute
setInterval(updateTimestamps, 60000);

// Initialize timestamps when the page loads
document.addEventListener('DOMContentLoaded', updateTimestamps);