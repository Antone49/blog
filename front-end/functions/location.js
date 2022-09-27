import {sendRequest} from './broker.js';

async function getAllLocations() {
  console.log('getAllLocations');

  const payload = {
      action: 'getAllLocations',
  }

  return sendRequest(payload)
}

async function removeLocation(token, id) {
    console.log('removeLocation');

    const payload = {
        action: 'removeLocation',
        token: token,
        location: {
            id: id
        }
    }

    return sendRequest(payload)
}

async function getLocation(token, id) {
    console.log('getLocation');

    const payload = {
        action: 'getLocation',
        token: token,
        location: {
            id: id
        }
    }

    return sendRequest(payload)
}

async function addLocation(token, name, image, longitude, latitude) {
    console.log('addLocation');

    const payload = {
        action: 'addLocation',
        token: token,
        location: {
            name: name,
            image: image,
            longitude: longitude,
            latitude: latitude,
        }
    }

    return sendRequest(payload)
}

async function updateLocation(token, id, name, image, longitude, latitude) {
    console.log('updateLocation');

    const payload = {
        action: 'updateLocation',
        token: token,
        location: {
            id: id,
            name: name,
            image: image,
            longitude: longitude,
            latitude: latitude,
        }
    }

    return sendRequest(payload)
}

export { getAllLocations, getLocation, removeLocation, addLocation, updateLocation }