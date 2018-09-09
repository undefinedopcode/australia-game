// The latitude and longitude of your business / place
var position = [27.1959739, 78.02423269999997];
var ne = [27.1959739, 78.02423269999997];
var sw = [27.1959739, 78.02423269999997];
var showBox = false;

function updateMap( lat, lon, nelat, nelon, swlat, swlon ) {
    position = [Number(lat), Number(lon)];
    ne = [Number(nelat), Number(nelon)];
    sw = [Number(swlat), Number(swlon)];
    showGoogleMaps();
}

function showGoogleMaps() {

    var latLng = new google.maps.LatLng(position[0], position[1]);
    var nell = new google.maps.LatLng(ne[0], ne[1]);
    var swll = new google.maps.LatLng(sw[0], sw[1]);

    var bounds = new google.maps.LatLngBounds(swll, nell);

    var mapOptions = {
        zoom: 15, // initialize zoom level - the max value is 21
        mapTypeId: google.maps.MapTypeId.SATELLITE,
        disableDefaultUI: true,
        center: latLng,
        bounds: bounds,
    };

    map = new google.maps.Map(document.getElementById('map-background'),
        mapOptions);
    map.fitBounds(bounds);

    // Show the default red marker at the location
    marker = new google.maps.Marker({
        position: latLng,
        map: map,
        draggable: false,
        animation: google.maps.Animation.DROP
    });

    if (showBox) {
        rectangle = new google.maps.Rectangle({
            strokeColor: '#FF0000',
            strokeOpacity: 0.8,
            strokeWeight: 2,
            fillColor: '#FF0000',
            fillOpacity: 0.35,
            map: map,
            bounds: {
            north: ne[0],
            south: sw[0],
            east:  ne[1],
            west: sw[1],
            }
        });
    }
}

google.maps.event.addDomListener(window, 'load', showGoogleMaps);