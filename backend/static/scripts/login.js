import { AuthService } from "./authservice.js";

document.addEventListener("DOMContentLoaded", function () {
	// Initialize Lucide icons
	lucide.createIcons();
	// DOM Elements
	const form = document.getElementById("signinForm");
	const emailInput = document.getElementById("email");
	const passwordInput = document.getElementById("password");
	const passwordToggle = document.getElementById("passwordToggle");
	const submitButton = form.querySelector('button[type="submit"]');
	const spinner = submitButton.querySelector(".spinner");

	// Initialize AuthService
	const authService = new AuthService();

	// Password visibility toggle
	passwordToggle.addEventListener("click", () => {
		const type = passwordInput.type === "password" ? "text" : "password";
		passwordInput.type = type;
		// Update icon
		const icon = passwordToggle.querySelector("i");
		if (icon != null) {
			icon.setAttribute("data-lucide", type === "password" ? "eye" : "eye-off");
		}
	});

	// Form validation
	function validateEmail(email) {
		return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
	}

	function showError(input, message) {
		const formGroup = input.closest(".form-group");
		let errorDiv = formGroup.querySelector(".error");

		if (!errorDiv) {
			errorDiv = document.createElement("div");
			errorDiv.className = "error";
			formGroup.appendChild(errorDiv);
		}

		errorDiv.textContent = message;
		input.classList.add("error-input");
	}

	function removeError(input) {
		const formGroup = input.closest(".form-group");
		const error = formGroup.querySelector(".error");
		if (error) error.remove();
		input.classList.remove("error-input");
	}

	// Form submission
	form.addEventListener("submit", handleLogin);

	async function handleLogin(e) {
		e.preventDefault();

		removeError(emailInput);
		removeError(passwordInput);

		if (!validateEmail(emailInput.value)) {
			showError(emailInput, "Please enter a valid email address");
			return;
		}

		if (passwordInput.value.length < window.CONSTANTS.MIN_PASSWORD_LENGTH) {
			showError(
				passwordInput,
				`Password must be at least ${window.CONSTANTS.MIN_PASSWORD_LENGTH} characters`
			);
			return;
		}

		// Show loading state
		submitButton.disabled = true;
		spinner.classList.remove("hidden");
		submitButton.querySelector("span").textContent = "Signing in...";

		const formData = {
			email: emailInput.value,
			password: passwordInput.value,
		};

		const response = await authService.login(formData);

		if (!response.error) {
			console.log("Login successful:", response.message);
			localStorage.setItem("res", JSON.stringify(response.data));
			window.location.href = "/";
		} else {
			console.error("Login failed:", response.message || "Unknown error");
			showError(emailInput, response.message || "Invalid email or password");
		}

		submitButton.disabled = false;
		spinner.classList.add("hidden");
		submitButton.querySelector("span").textContent = "Sign in";
	}

	// Clear errors dynamically
	form.addEventListener("input", (event) => {
		if (event.target.matches("#email, #password")) {
			removeError(event.target);
		}
	});

	// Reset profile data
	if (window.RESDATA.userData) {
		localStorage.removeItem("res");
		window.RESDATA.userData = null;
		window.RESDATA.profileImageElement.src = "";
		window.RESDATA.profileImageElement.alt = "";
		window.location.reload();
	}
});
