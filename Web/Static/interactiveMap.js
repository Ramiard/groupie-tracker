window.addEventListener('DOMContentLoaded', () => {
    // Setting up the mapbox accessToken
    mapboxgl.accessToken = 'pk.eyJ1IjoicmFtaWFyZDEyIiwiYSI6ImNtNXBibmE1cTA4bWcybXNpaWg4cWgydDgifQ.eZ661gjAiBWtfYMxXvN9Hw';



    // Setting up the map style
    const map = new mapboxgl.Map({
        // style: 'mapbox://styles/mapbox/outdoors-v12', // mapbox style URL
        style: 'mapbox://styles/ramiard12/cm5y7wkf1000301s9gzem8rp7', // our style URL
        container: 'map', // container ID
        center: [-33.002872,30.973201], // starting position [lng, lat]. Note that lat must be set between -90 and 90
        pitch: 60, // The angle of the map in degrees (0 is straight down)
        bearing: 0,
        zoom: 1, // starting zoom
        projection: 'mercator'
    });

    // Adding the markers one by one to the map
    let i =0;
    for (const [location , value] of Object.entries(relation)){
    marker = new mapboxgl.Marker({ color: 'red' })
        .setLngLat(concertCoordinates[i])
        .setPopup(new mapboxgl.Popup().setHTML(`<h2>${location}</h2> <ul>${value.map(date => `<li>${date}</li>`).join('')}</ul>`))
        .addTo(map);
    i++;
    }

    // Jump to all the locations
    map.on('load', async () => {

        // Set currentPopup to null
        let currentPopup = null;


        for (const coordinate of concertCoordinates) {
            // Initialize the city and dates variables
            // 'cityName' is the key that correspond to the current index of 'coordinate' in the 'relations' map
            const cityName = Object.keys(relation)[concertCoordinates.indexOf(coordinate)];
            // 'dates' is an array of all the dates that correspond to the current 'cityName'
            const dates = relation[cityName];

            // If there is a popup, remove it before displaying the new one
            if (currentPopup) {
                currentPopup.remove();
                currentPopup = null;
            }

            map.flyTo({ center: coordinate, zoom: 12 });
            map.once('moveend', () => {
                // Create and display popup after flyTo completes
                currentPopup = new mapboxgl.Popup({ offset: 25 })
                    .setLngLat(coordinate)
                    // Display the current city name and do a range on the dates array to display them in a list
                    .setHTML(`<h3>${cityName}</h3><ul>${dates.map(date => `<li>${date}</li>`).join('')}</ul>`)
                    .addTo(map)
            });
            // Wait 10 seconds before moving to the next location
            await new Promise(resolve => setTimeout(resolve, 10000));
        }
    });
});
