const sidebar = document.querySelector('#sidebar'); 

// Handle window resize 
function handleResize() { 
    if (window.innerWidth >= 768) { 
        sidebar.classList.add('visible'); 
        sidebar.classList.remove('hidden'); 
    } else { 
        sidebar.classList.add('hidden'); 
        sidebar.classList.remove('visible'); 
    }
}

// Initialize function 
function init() {
    // Initial resize check 
    handleResize(); 
}

// Start the application 
document.addEventListener('DOMContentLoaded', init);

export { handleResize, sidebar};