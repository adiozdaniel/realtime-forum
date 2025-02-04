document.addEventListener('DOMContentLoaded', () => { 
    lucide.createIcons();
    const modal = document.getElementById('modal'); 
    const createPostBtn = document.getElementById('createPostBtn'); 
    const cancelBtn = document.getElementById('cancelBtn'); 
    const postForm = document.getElementById('postForm'); 

    // Open modal 
    createPostBtn.addEventListener('click', () => { 
        modal.classList.add('open'); 
    }); 

    // Close modal when clicking outside 
    modal.addEventListener('click', (e) => { 
        if (e.target === modal) { 
            modal.classList.remove('open'); 
        } 
    }); 

    // Close modal with cancel button 
    cancelBtn.addEventListener('click', () => { 
        modal.classList.remove('open'); 
    }); 

    // Handle form submission 
    postForm.addEventListener('submit', (e) => { 
        e.preventDefault(); const formData = { 
            title: document.getElementById('title').value, 
            author: document.getElementById('author').value, 
            category: document.getElementById('category').value, 
            excerpt: document.getElementById('excerpt').value 
        };

        // Send the data to a server 
        console.log('Form submitted:', formData);
         
        // Reset form and close modal 
        postForm.reset(); modal.classList.remove('open'); 
    }); 
});