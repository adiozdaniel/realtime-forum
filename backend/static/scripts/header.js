// DOM Elements 
const darkModeToggle = document.querySelector('#darkModeToggle'); 
const authButton = document.querySelector(".sign-in-button");


// Toggle dark mode 
function toggleDarkMode() { 
    document.body.classList.toggle('dark-mode'); localStorage.setItem('darkMode', document.body.classList.contains('dark-mode')); 
}

async function signOutUser() {
    try {
        let response = await fetch("/api/auth/logout", {
            method: "POST", 
            credentials: "include" 
        });

        if (!response.ok) console.log("Not logged in");

        if (response.ok) {
            console.log("User signed out successfully.");
        }

        return response.status === 200;
    } catch (error) {
        console.error("Error signing out:", error);
    }
}

async function isSignedIn() {
    if (!authButton) {
        console.error("Auth button not found.");
        return;
    }

    try {
        let response = await fetch("/api/auth/check", { credentials: "include" });
        if (!response.ok) {
            throw new Error("Not signed in")
        };

        let data = await response.json();

        return data.signedIn;
    } catch (error) {
        return false;
    }
}

// Initialize function 
function init() {
    lucide.createIcons();

    authButton.textContent = isSignedIn() ? "Sign Out" : "Sign In";

    // Automatically log out if on /auth
    if (window.location.pathname === "/auth") {
        signOutUser();
        authButton.textContent = "Sign In";
    };

    // Event listeners 
    darkModeToggle?.addEventListener('click', toggleDarkMode); 
    // Check for saved dark mode preference 
    const savedDarkMode = localStorage.getItem('darkMode') === 'true'; 
    if (savedDarkMode) { document.body.classList.add('dark-mode'); } 
} 

// Start the application 
document.addEventListener('DOMContentLoaded', init);