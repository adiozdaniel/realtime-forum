import { PostService } from './postsservice.js';

class PostModalManager {
    constructor() {}
    modal = document.getElementById('createPostModal');
    createPostBtn = document.getElementById('createPostBtn');
    cancelBtn = document.getElementById('cancelPost');
    form = document.getElementById('createPostForm');
    imageUpload = document.getElementById('imageUpload');
    imagePreviewContainer = document.getElementById('imagePreviewContainer');
    imagePreview = document.getElementById('imagePreview');
    removeImage = document.getElementById('removeImage');
    videoLink = document.getElementById('videoLink');
    videoPreviewContainer = document.getElementById('videoPreviewContainer');
    videoPreview = document.getElementById('videoPreview');
    removeVideo = document.getElementById('removeVideo');
    mediaPreview = document.getElementById('mediaPreview');
    uploadError = document.getElementById('uploadError');
    MAX_FILE_SIZE = 20 * 1024 * 1024;
    ALLOWED_TYPES = ['image/jpeg', 'image/png', 'image/gif'];
    postService = new PostService();
}

PostModalManager.prototype.init = function () {
    lucide.createIcons();
    this.createPostBtn.addEventListener('click', this.openModal.bind(this));
    this.cancelBtn.addEventListener('click', this.closeModal.bind(this));
    this.modal.addEventListener('click', (e) => {
        if (e.target === this.modal) this.closeModal();
    });
    this.imageUpload.addEventListener('change', this.handleImageUpload.bind(this));
    this.videoLink.addEventListener('input', this.handleVideoLink.bind(this));
    this.removeImage.addEventListener('click', this.removeImagePreview.bind(this));
    this.removeVideo.addEventListener('click', this.removeVideoPreview.bind(this));
    this.form.addEventListener('submit', this.handleSubmit.bind(this));
};

PostModalManager.prototype.openModal = function () {
    this.modal.classList.remove('hidden');
};

PostModalManager.prototype.closeModal = function () {
    this.modal.classList.add('hidden');
    this.form.reset();
    this.imagePreview.src = '';
    this.imagePreviewContainer.classList.add('hidden');
    this.videoPreviewContainer.classList.add('hidden');
    this.mediaPreview.classList.add('hidden');
    this.uploadError.classList.add('hidden');
    this.uploadError.textContent = '';
};

PostModalManager.prototype.handleImageUpload = function (e) {
    const file = e.target.files[0];
    this.uploadError.textContent = '';
    this.uploadError.classList.add('hidden');

    if (!file) return;

    if (!this.ALLOWED_TYPES.includes(file.type)) {
        this.showUploadError('Invalid file type. Please upload a JPEG, PNG, or GIF image.');
        this.imageUpload.value = '';
        return;
    }

    if (file.size > this.MAX_FILE_SIZE) {
        this.showUploadError('File size exceeds 20MB limit.');
        this.imageUpload.value = '';
        return;
    }

    const reader = new FileReader();
    reader.onload = (e) => {
        this.imagePreview.src = e.target.result;
        this.imagePreviewContainer.classList.remove('hidden');
        this.mediaPreview.classList.remove('hidden');
        this.videoLink.value = '';
        this.videoPreviewContainer.classList.add('hidden');
    };
    reader.onerror = () => {
        this.showUploadError('Error reading file. Please try again.');
        this.imageUpload.value = '';
    };
    reader.readAsDataURL(file);
};

PostModalManager.prototype.handleVideoLink = function (e) {
    const url = e.target.value.trim();
    if (this.isValidVideoUrl(url)) {
        const embedUrl = this.getEmbedUrl(url);
        if (embedUrl) {
            this.videoPreview.innerHTML = `<iframe width="100%" height="250" src="${embedUrl}" frameborder="0" allowfullscreen></iframe>`;
            this.videoPreviewContainer.classList.remove('hidden');
            this.mediaPreview.classList.remove('hidden');
            this.imageUpload.value = '';
            this.imagePreviewContainer.classList.add('hidden');
        }
    } else {
        this.videoPreviewContainer.classList.add('hidden');
    }
};

PostModalManager.prototype.removeImagePreview = function () {
    this.imageUpload.value = '';
    this.imagePreview.src = '';
    this.imagePreviewContainer.classList.add('hidden');
    if (this.videoPreviewContainer.classList.contains('hidden')) {
        this.mediaPreview.classList.add('hidden');
    }
};

PostModalManager.prototype.removeVideoPreview = function () {
    this.videoLink.value = '';
    this.videoPreviewContainer.classList.add('hidden');
    if (this.imagePreviewContainer.classList.contains('hidden')) {
        this.mediaPreview.classList.add('hidden');
    }
};

PostModalManager.prototype.handleSubmit = async function (e) {
    e.preventDefault();
    const formData = {
        PostTitle: document.getElementById('postTitle').value,
        PostCategory: document.getElementById('postCategory').value,
        PostContent: document.getElementById('postContent').value,
        PostImage: null,
        PostVideo: this.videoLink.value || null
    };
    const imageFile = this.imageUpload.files[0];
    if (imageFile) {
        formData.image = {
            name: imageFile.name,
            type: imageFile.type,
            size: imageFile.size
        };
    }
    const res = await this.postService.createPost(formData);
    console.log(res);
    if (!res.error) {
        this.closeModal();
    }
};

PostModalManager.prototype.showUploadError = function (message) {
    this.uploadError.textContent = message;
    this.uploadError.classList.remove('hidden');
    this.uploadError.classList.remove('error-shake');
    void this.uploadError.offsetWidth;
    this.uploadError.classList.add('error-shake');
};

PostModalManager.prototype.isValidVideoUrl = function (url) {
    return /^(https?:\/\/)?(www\.)?(youtube\.com|youtu\.be|vimeo\.com)\/.+/.test(url);
};

PostModalManager.prototype.getEmbedUrl = function (url) {
    const youtubeRegex = /(?:youtube\.com\/(?:watch\?v=|embed\/|shorts\/)|youtu\.be\/|youtube\.com\/live\/)([a-zA-Z0-9_-]+)/;
    const vimeoRegex = /vimeo\.com\/(\d+)/;
    let match = url.match(youtubeRegex);
    if (match) return `https://www.youtube.com/embed/${match[1]}`;
    match = url.match(vimeoRegex);
    if (match) return `https://player.vimeo.com/video/${match[1]}`;
    return null;
};

export {PostModalManager};
