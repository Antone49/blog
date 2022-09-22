import {sendRequest} from './broker.js';

async function getAllPosts(search) {
  console.log('getAllPosts');

  const payload = {
      action: 'getAllPosts',
      post: {
        search: search
      }
  }

  return sendRequest(payload)
}

async function getPost(id) {
    console.log('getPost');
  
    const payload = {
        action: 'getPost',
        post: {
            id: id
        }
    }
  
    return sendRequest(payload)
}

async function getAllTags() {
    console.log('getAllTags');

    const payload = {
        action: 'getAllTags',
    }

    return sendRequest(payload)
}

async function removeTag(token, name) {
    console.log('removeTag');

    const payload = {
        action: 'removeTag',
        token: token,
        tag: {
            id: name
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
            id: name
        }
    }

    return sendRequest(payload)
}

async function updateTag(token, oldName, newName) {
    console.log('updateTag');

    const payload = {
        action: 'updateTag',
        token: token,
        tag: {
            oldName: oldName,
            newName: newName
        }
    }

    return sendRequest(payload)
}

async function getPostTags(id) {
    console.log('getPostTags');

    const payload = {
        action: 'getPostTags',
        post: {
            id: id
        }
    }

    return sendRequest(payload)
}
  

async function getLastestPosts(number) {
    console.log('getLastestPosts');

    const payload = {
        action: 'getLastestPosts',
        post: {
        number: number
        }
    }

    return sendRequest(payload)
}

async function getAllLocations() {
  console.log('getAllLocations');

  const payload = {
      action: 'getAllLocations',
  }

  return sendRequest(payload)
}

export { getAllPosts, getPost, getAllTags, addTag, updateTag, removeTag, getPostTags, getLastestPosts, getAllLocations }