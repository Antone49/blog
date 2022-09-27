import {sendRequest} from './broker.js';

async function getAllTags() {
    console.log('getAllTags');

    const payload = {
        action: 'getAllTags',
    }

    return sendRequest(payload)
}

async function removeTag(token, id) {
    console.log('removeTag');

    const payload = {
        action: 'removeTag',
        token: token,
        tag: {
            id: id
        }
    }

    return sendRequest(payload)
}

async function addTag(token, name) {
    console.log('addTag');

    const payload = {
        action: 'addTag',
        token: token,
        tag: {
            name: name
        }
    }

    return sendRequest(payload)
}

async function updateTag(token, id, name) {
    console.log('updateTag');

    const payload = {
        action: 'updateTag',
        token: token,
        tag: {
            id: id,
            name: name
        }
    }

    return sendRequest(payload)
}

export { getAllTags, addTag, updateTag, removeTag }