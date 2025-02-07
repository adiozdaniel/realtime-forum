// DOM Elements 
const darkModeToggle = document.querySelector('#darkModeToggle'); 
const authButton = document.querySelector(".sign-in-button");


// Toggle dark mode 
function toggleDarkMode() { 
    document.body.classList.toggle('dark-mode'); localStorage.setItem('darkMode', document.body.classList.contains('dark-mode')); 
}

async function isSignedIn() {
    if (!authButton) {
        console.error("Auth button not found.");
        return;
    }

    try {
        let response = await fetch("/api/auth/check", { credentials: "include" });
        if (!response.ok) {
            authButton.textContent = "Sign In";
            throw new Error("Not signed in")
        };
        authButton.textContent = "Sign Out";

        let data = await response.json();
        return data.signedIn;
    } catch (error) {
        console.log(error);
        return false;
    }
}

// Initialize function 
function init() {
    lucide.createIcons();
    // Event listeners 
    darkModeToggle?.addEventListener('click', toggleDarkMode); 
    // Check for saved dark mode preference 
    const savedDarkMode = localStorage.getItem('darkMode') === 'true'; 
    if (savedDarkMode) { document.body.classList.add('dark-mode'); } 

    isSignedIn().then(signedIn => console.log("User signed in:", signedIn));
} 

// Start the application 
document.addEventListener('DOMContentLoaded', init);