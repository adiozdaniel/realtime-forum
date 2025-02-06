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

    // Password visibility toggle 
    passwordToggle.addEventListener('click', () => { 
        const type = passwordInput.type === 'password' ? 'text' : 'password'; 
        passwordInput.type = type; 
        confirmPasswordInput.type = type; 
        // Update icon 
        const icon = passwordToggle.querySelector('i'); 
        icon.setAttribute('data-lucide', type === 'password' ? 'eye' : 'eye-off'); 
        lucide.createIcons(); 
    }); 

     // Check password match
     function checkPasswordsMatch() { 
        if (confirmPasswordInput.value && passwordInput.value !== confirmPasswordInput.value) { 
            showError(confirmPasswordInput, 'Passwords do not match'); 
            return false; 
        } else { 
            removeError(confirmPasswordInput); 
            return true; 
        } 
    } 

    // Password strength checker 
    function checkPasswordStrength(password) { 
        let strength = 0; 
        if (password.length >= 8) strength++; 
        if (password.match(/[a-z]/) && password.match(/[A-Z]/)) strength++; 
        if (password.match(/\d/)) strength++; if (password.match(/[^a-zA-Z\d]/)) strength++; 

        // Update strength bars 
        strengthBars.forEach((bar, index) => { 
            bar.className = 'bar'; 
            if (index < strength) { 
                bar.classList.add(strength <= 2 ? 'weak' : strength === 3 ? 'medium' : 'strong'); 
            } 
        }); 

        // Update strength text 
        const strengthLabels = ['Weak', 'Fair', 'Good', 'Strong']; 
        strengthText.textContent = strength > 0 ? strengthLabels[strength - 1] : 'Password strength'; 
        return strength; 
    } 

    // Form validation 
    function validateEmail(email) { 
        return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email); 
    } 

    function showError(input, message) { 
        const formGroup = input.closest('.form-group'); 
        const existingError = formGroup.querySelector('.error'); 
        if (!existingError) { 
            const errorDiv = document.createElement('div'); 
            errorDiv.className = 'error'; 
            errorDiv.textContent = message; 
            formGroup.appendChild(errorDiv); 
        } 
        input.classList.add('error-input'); 
    } 

    function removeError(input) { 
        const formGroup = input.closest('.form-group'); 
        const error = formGroup.querySelector('.error'); 
        if (error) { error.remove(); 

        } 
        input.classList.remove('error-input'); 
    } 

    // Input event listeners 
    passwordInput.addEventListener('input', () => { 
        removeError(passwordInput); 
        checkPasswordStrength(passwordInput.value); 
        if (confirmPasswordInput.value) { 
            checkPasswordsMatch(); 
        } 
    }); 

    confirmPasswordInput.addEventListener('input', () => { 
       checkPasswordsMatch();
    }); 

    // Form submission 
    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        let isValid = true;
    
        // Clear previous errors
        [usernameInput, emailInput, passwordInput, confirmPasswordInput].forEach(removeError);
    
        // Validate username
        if (usernameInput.value.length < 3) {
            showError(usernameInput, 'Username must be at least 3 characters');
            isValid = false;
        }
    
        // Validate email
        if (!validateEmail(emailInput.value)) {
            showError(emailInput, 'Please enter a valid email address');
            isValid = false;
        }
    
        // Validate password
        const passwordStrength = checkPasswordStrength(passwordInput.value);
        if (passwordStrength < 3) {
            showError(passwordInput, 'Password is too weak');
            isValid = false;
        }
    
        // Validate password confirmation
        if (!checkPasswordsMatch()) {
            isValid = false;
        }
    
        // Validate terms
        if (!termsCheckbox.checked) {
            showError(termsCheckbox, 'You must accept the terms and conditions');
            isValid = false;
        }
    
        if (!isValid) return;
    
        // Show loading state
        submitButton.disabled = true;
        spinner.classList.remove('hidden');
        submitButton.querySelector('span').textContent = 'Creating account...';
    
        // Prepare form data
        const formData = {
            email: emailInput.value,
            username: usernameInput.value,
            password: passwordInput.value,
        };
    
        try {
            const res = await fetch('/api/auth/register', {
                method: 'POST',
                body: JSON.stringify(formData),
                headers: {
                    'Content-Type': 'application/json',
                },
            });
    
            // Log the response status
            console.log('Response Status:', res.status);
    
            // Parse the response JSON
            const response = await res.json();
    
            // Log the response data
            console.log('Response Data:', response);
    
            // Store the response in localStorage
            localStorage.setItem('res', JSON.stringify(response));
    
            // Handle success or error based on the response
            if (res.ok) {
                console.log('Registration successful:', response.message);
                // Redirect or show a success message
                window.location.href = '/';
            } else {
                console.error('Registration failed:', response.message || 'Unknown error');
                showError(emailInput, response.message || 'Failed to create account. Please try again.');
            }
        } catch (error) {
            // Log any network or unexpected errors
            console.error('Error during registration:', error);
            showError(emailInput, 'Failed to create account. Please try again.');
        } finally {
            // Reset button state
            submitButton.disabled = false;
            spinner.classList.add('hidden');
            submitButton.querySelector('span').textContent = 'Create account';
        }
    });
});
