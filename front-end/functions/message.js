import {sendRequest} from './broker.js';

async function addMessage(username, comment, email, postId) {
  console.log('addMessage');

  const payload = {
      action: 'addMessage',
      message: {
        username: username,
        message: comment,
        email: email,
        postId: postId
      }
  }

  return sendRequest(payload)
}

async function getAllMessagesFromPostId(postId) {
  console.log('getAllMessagesFromPostId');

  const payload = {
      action: 'getAllMessagesFromPostId',
      message: {
        postId: postId
      }
  }

  return sendRequest(payload)
}

async function getAllMessages(token) {
  console.log('getAllMessages');

  const payload = {
      action: 'getAllMessages',
      token: token,
  }

  return sendRequest(payload)
}

async function updateValidityMessage(token, id, valid) {
  console.log('updateValidityMessage');

  const payload = {
      action: 'updateValidityMessage',
      token: token,
      message: {
        id: id,
        valid: valid,
      }
  }

  return sendRequest(payload)
}

async function removeMessage(token, id) {
  console.log('removeMessage');

  const payload = {
      action: 'removeMessage',
      token: token,
      message: {
        id: id,
      }
  }

  return sendRequest(payload)
}

export { addMessage, getAllMessagesFromPostId, getAllMessages, updateValidityMessage, removeMessage }