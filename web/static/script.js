const BACKEND_URL = "https://0.0.0.0:5000/"
const urlForm = document.getElementById("urlForm")

urlForm.addEventListener("submit", event => {
	event.preventDefault();
	const form = event.currentTarget;
	console.log(form);
	const formData = new FormData(form);
	const jsonData = Object.fromEntries(formData.entries());
	const requestData = JSON.stringify(jsonData);
	const requestOptions = {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
			"Accept": "application/json"
		},
		body: requestData,
	};
	console.log(requestData);
	const response = fetch(BACKEND_URL, requestOptions);
	if (!response.ok) {
		console.error(response.text);
	}
	const responseData = response.json;
	console.log(responseData);
});
