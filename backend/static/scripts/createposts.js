document.addEventListener('DOMContentLoaded', () => { 
    
    // Initialize Lucide icons 
    lucide.createIcons();

    // Elements 
    const modal = document.getElementById('createPostModal'); 
    const createPostBtn = document.getElementById('createPostBtn'); 
    const cancelBtn = document.getElementById('cancelPost'); 
    const form = document.getElementById('createPostForm'); 
    const imageUpload = document.getElementById('imageUpload'); 
    const imagePreviewContainer = document.getElementById('imagePreviewContainer'); 
    const imagePreview = document.getElementById('imagePreview'); 
    const removeImage = document.getElementById('removeImage'); 
    const videoLink = document.getElementById('videoLink'); 
    const videoPreviewContainer = document.getElementById('videoPreviewContainer'); 
    const videoPreview = document.getElementById('videoPreview'); 
    const removeVideo = document.getElementById('removeVideo'); 
    const mediaPreview = document.getElementById('mediaPreview');
    const MAX_FILE_SIZE = 20 * 1024 * 1024; // 20MB in bytes 
    const ALLOWED_TYPES = ['image/jpeg', 'image/png', 'image/gif']; 
    const uploadError = document.getElementById('uploadError'); 
    
    // Open modal 
    createPostBtn.addEventListener('click', () => { 
        modal.classList.remove('hidden'); 
    }); 
    
    // Close modal 
    cancelBtn.addEventListener('click', closeModal); 
    modal.addEventListener('click', (e) => { 
        if (e.target === modal) closeModal(); 
    }); 
    
    // Handle image upload 
    imageUpload.addEventListener('change', (e) => { 
        const file = e.target.files[0]; 
        
        // Reset error state 
        uploadError.textContent = ''; 
        uploadError.classList.add('hidden'); 

        if (!file) return; 

        // Validate file type 
        if (!ALLOWED_TYPES.includes(file.type)) { 
            showUploadError('Invalid file type. Please upload a JPEG, PNG, or GIF image.'); 
            imageUpload.value = ''; 
            return; 
        } 
        
        // Validate file size 
        if (file.size > MAX_FILE_SIZE) { 
            showUploadError('File size exceeds 20MB limit.'); 
            imageUpload.value = ''; 
            return; 
        } 
        
        // If validation passes, show preview 
        const reader = new FileReader(); 
        reader.onload = (e) => { 
            imagePreview.src = e.target.result; 
            imagePreviewContainer.classList.remove('hidden'); 
            mediaPreview.classList.remove('hidden');
                
                // Clear video if exists 
                videoLink.value = ''; 
                videoPreviewContainer.classList.add('hidden'); 
        }; 

        reader.onerror = () => { 
            showUploadError('Error reading file. Please try again.'); 
            imageUpload.value = ''; 
        }; 
            
        reader.readAsDataURL(file); 
        
    }); 
    
    function showUploadError(message) { 
        uploadError.textContent = message; 
        uploadError.classList.remove('hidden'); 
        uploadError.classList.remove('error-shake'); 
        
        // Trigger reflow to restart animation void 
        uploadError.offsetWidth; uploadError.classList.add('error-shake');
    }
    
    // Handle video link 
    videoLink.addEventListener('input', (e) => { 
        const url = e.target.value.trim(); 
        if (isValidVideoUrl(url)) { 
            const embedUrl = getEmbedUrl(url);
        
        if (embedUrl) {
            videoPreview.innerHTML = `<iframe width="100%" height="250" src="${embedUrl}" frameborder="0" allowfullscreen></iframe>`;
            videoPreviewContainer.classList.remove('hidden'); 
            mediaPreview.classList.remove('hidden'); 
            
            // Clear image if exists 
            imageUpload.value = ''; 
            imagePreviewContainer.classList.add('hidden'); 
        }
        } else { 
            videoPreviewContainer.classList.add('hidden'); 
        } 
    }); 
    
    // Remove image 
    removeImage.addEventListener('click', () => { 
        imageUpload.value = ''; imagePreview.src = ''; 
        imagePreviewContainer.classList.add('hidden'); 
        if (videoPreviewContainer.classList.contains('hidden')) { 
            mediaPreview.classList.add('hidden'); 
        } 
    }); 
    
    // Remove video 
    removeVideo.addEventListener('click', () => { 
        videoLink.value = ''; 
        videoPreviewContainer.classList.add('hidden'); 
        if (imagePreviewContainer.classList.contains('hidden')) { 
            mediaPreview.classList.add('hidden'); 
        } 
    }); 

    // Handle form submission 
    form.addEventListener('submit', (e) => { 
        e.preventDefault(); 
        const formData = { 
            title: document.getElementById('postTitle').value, 
            author: document.getElementById('postAuthor').value, 
            category: document.getElementById('postCategory').value, 
            content: document.getElementById('postContent').value, 
            image: null, 
            video: videoLink.value || null 
        }; 

        // Handle image data 
        const imageFile = imageUpload.files[0]; 
        if (imageFile) { 
            // Here you would typically handle the file upload to your server // For now, we'll just add the file information 
            formData.image = { 
                name: imageFile.name, 
                type: imageFile.type, 
                size: imageFile.size 
            }; 
        } 

        console.log('Post Data:', formData); 

        closeModal(); 
    }); 
    
    function closeModal() { 
        modal.classList.add('hidden'); 
        form.reset(); 
        imagePreview.src = ''; 
        imagePreviewContainer.classList.add('hidden'); 
        videoPreviewContainer.classList.add('hidden'); 
        mediaPreview.classList.add('hidden'); 
        uploadError.classList.add('hidden'); 
        uploadError.textContent = ''; 
    }
        
    function isValidVideoUrl(url) { 
        // Basic validation for YouTube and Vimeo URLs 
        return /^(https?:\/\/)?(www\.)?(youtube\.com|youtu\.be|vimeo\.com)\/.+/.test(url); 
    } 

    function getEmbedUrl(url) {
        const youtubeRegex = /(?:youtube\.com\/(?:watch\?v=|embed\/|shorts\/)|youtu\.be\/|youtube\.com\/live\/)([a-zA-Z0-9_-]+)/;
        const vimeoRegex = /vimeo\.com\/(\d+)/;
    
        let match = url.match(youtubeRegex);
        if (match) {
            return `https://www.youtube.com/embed/${match[1]}`;
        }
    
        match = url.match(vimeoRegex);
        if (match) {
            return `https://player.vimeo.com/video/${match[1]}`;
        }
    
        return null; // Invalid video link
    }
    
    
});