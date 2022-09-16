import {sendRequest} from './broker.js';

async function getAllPosts() {
  console.log('getAllPosts');

  const payload = {
      action: 'getAllPosts',
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

export { getAllPosts, getPost, getPostTags, getLastestPosts }