export const views = {
	home: () => `
        <h1>Home</h1>
        <p>Welcome to our SPA.</p>
        <a href="/about" data-link>Go to About</a>
    `,
	about: () => `
        <h1>About</h1>
        <p>Learn more about our application.</p>
        <a href="/" data-link>Go to Home</a>
    `,
	contact: () => `
        <h1>Contact</h1>
        <p>Get in touch with us.</p>
        <a href="/" data-link>Go to Home</a>
    `,
	notFound: () => `
        <h1>404 - Page Not Found</h1>
        <p>The page you are looking for does not exist.</p>
        <a href="/" data-link>Go to Home</a>
    `,
};
