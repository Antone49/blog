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

async function addPost(token) {
    console.log('addPost');

    const payload = {
        action: 'addPost',
        token: token,
    }

    return sendRequest(payload)
}

async function updatePost(token, id, title, content) {
    console.log('updatePost');

    const payload = {
        action: 'updatePost',
        token: token,
        post: {
            id: id,
            title: title,
            content: content,
        }
    }

    return sendRequest(payload)
}

async function updatePostImage(token, id, image) {
    console.log('updatePostImage');

    const payload = {
        action: 'updatePostImage',
        token: token,
        post: {
            id: id,
            image: image,
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

async function getAllPostLocations() {
    console.log('getAllPostLocations');

    const payload = {
        action: 'getAllPostLocations',
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

export { getAllPosts, getPost, addPost, removePost, updatePost, updatePostImage, getPostTags, updatePostTags, getLastestPosts, getAllPostLocations, getPostLocations, updatePostLocations }