import {sendRequest} from './broker.js';

async function getAllPosts(search) {
  console.log('getAllPosts');

  const payload = {
      action: 'getAllPosts',
      token: 'CGfj9PPdZ6WJwrCVvH6Z4Oi22S07TTEmziT_QTsr2NA',
      allPosts: {
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

export { getAllPosts, getPost, getAllTags, getPostTags, getLastestPosts, getAllLocations }