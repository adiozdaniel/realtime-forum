document.addEventListener('DOMContentLoaded', function () {
    // Initialize Lucide icons
    lucide.createIcons();

    // DOM Elements
    const form = document.getElementById('signinForm');
    const emailInput = document.getElementById('email');
    const passwordInput = document.getElementById('password');
    const passwordToggle = document.getElementById('passwordToggle');
    const submitButton = form.querySelector('button[type="submit"]');
    const spinner = submitButton.querySelector('.spinner');

    // Password visibility toggle
    passwordToggle.addEventListener('click', () => {
        const isPassword = passwordInput.type === 'password';
        passwordInput.type = isPassword ? 'text' : 'password';
        passwordToggle.querySelector('i').setAttribute('data-lucide', isPassword ? 'eye-off' : 'eye');
        lucide.createIcons(passwordToggle);
    });

    // Form validation
    function validateEmail(email) {
        return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
    }

    function showError(input, message) {
        const formGroup = input.closest('.form-group');
        let errorDiv = formGroup.querySelector('.error');

        if (!errorDiv) {
            errorDiv = document.createElement('div');
            errorDiv.className = 'error';
            formGroup.appendChild(errorDiv);
        }

        errorDiv.textContent = message;
        input.classList.add('error-input');
    }

    function removeError(input) {
        const formGroup = input.closest('.form-group');
        const error = formGroup.querySelector('.error');
        if (error) error.remove();
        input.classList.remove('error-input');
    }

    // Form submission with backend integration
    form.addEventListener('submit', async (e) => {
        e.preventDefault();

        removeError(emailInput);
        removeError(passwordInput);

        if (!validateEmail(emailInput.value)) {
            showError(emailInput, 'Please enter a valid email address');
            return;
        }

        if (passwordInput.value.length < 6) {
            showError(passwordInput, 'Password must be at least 6 characters');
            return;
        }

        // Show loading state
        submitButton.disabled = true;
        spinner.classList.remove('hidden');
        submitButton.querySelector('span').textContent = 'Signing in...';

        const formData = {
            email: emailInput.value,
            password: passwordInput.value,
        };

        let result;

        try {
            const response = await fetch('/api/auth/login', {
                method: 'POST',
                body: JSON.stringify(formData),
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            result = await response.json();
            console.log('Server Response:', result); // Log the response

            if (!response.ok) {
                throw new Error(result.message || 'Login failed');
            }

           // Redirect on success
            window.location.href = '/';
        } catch (error) {
            console.error('Login Error:', error.message);

            if (result.message === "user does not exist") {
                showError(emailInput, result.message);
                return;
            }
            
            showError(emailInput, result.message);
            showError(passwordInput, result.message);
        } finally {
            submitButton.disabled = false;
            spinner.classList.add('hidden');
            submitButton.querySelector('span').textContent = 'Sign in';
        }
    });

    // Clear errors dynamically
    form.addEventListener('input', (event) => {
        if (event.target.matches('#email, #password')) {
            removeError(event.target);
        }
    });
});
