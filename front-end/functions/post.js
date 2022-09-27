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

async function removePost(token, id) {
    console.log('removePost');

    const payload = {
        action: 'removePost',
        token: token,
        post: {
            id: id
        }
    }

    return sendRequest(payload)
}

async function addPost(token, title, image, content) {
    console.log('addPost');

    const payload = {
        action: 'addPost',
        token: token,
        post: {
            title: title,
            image: image,
            content: content,
        }
    }

    return sendRequest(payload)
}

async function updatePost(token, id, title, image, content) {
    console.log('updatePost');

    const payload = {
        action: 'updatePost',
        token: token,
        post: {
            id: id,
            title: title,
            image: image,
            content: content,
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

async function updatePostTags(token, postId, tagsId) {
    console.log('updatePostTags');

    const payload = {
        action: 'updatePostTags',
        token: token,
        postTag: {
            postId: postId,
            tagsId: tagsId
        }
    }

    return sendRequest(payload)
}

async function getPostLocations(id) {
    console.log('getPostLocations');

    const payload = {
        action: 'getPostLocations',
        post: {
            id: id
        }
    }

    return sendRequest(payload)
}

async function updatePostLocations(token, postId, locationId) {
    console.log('updatePostLocations');

    const payload = {
        action: 'updatePostLocations',
        token: token,
        postLocation: {
            postId: postId,
            locationId: locationId
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

export { getAllPosts, getPost, addPost, removePost, updatePost, getPostTags, updatePostTags, getLastestPosts, getPostLocations, updatePostLocations }