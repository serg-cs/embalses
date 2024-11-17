window.getLocation();

function getLocation() {
  if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(showPosition, showError);
  } else {
    console.log("Geolocation not supported by browser.")
	showLocationModal();
  }
}

function showPosition(position) {
	const lat = position.coords.latitude;
    const lon = position.coords.longitude;
	console.log("lat: " + lat + " lon: " + lon);

	const route = `/list/?lat=${lat}&lon=${lon}`;
	htmx.ajax('GET', route, {target:'#listView', swap:'innerHTML'})
}

function showError(error) {
	console.log("Error getting navigator geolocation.")
	console.log(error)
	showLocationModal();
}

function showLocationModal() {
	const locationModal = new bootstrap.Modal(document.getElementById('locationModal'), {keyboard: false});
	locationModal.show();
}
