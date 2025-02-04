document.addEventListener('DOMContentLoaded', function() { 
    // Initialize Lucide icons 
    lucide.createIcons(); 

    // DOM Elements 
    const form = document.getElementById('signupForm'); 
    const usernameInput = document.getElementById('username'); 
    const emailInput = document.getElementById('email'); 
    const passwordInput = document.getElementById('password'); 
    const confirmPasswordInput = document.getElementById('confirmPassword'); 
    const termsCheckbox = document.getElementById('terms'); 
    const passwordToggle = document.getElementById('passwordToggle'); 
    const strengthBars = document.querySelectorAll('.strength-bars .bar'); 
    const strengthText = document.querySelector('.strength-text'); 
    const submitButton = form.querySelector('button[type="submit"]'); 
    const spinner = submitButton.querySelector('.spinner'); 

});