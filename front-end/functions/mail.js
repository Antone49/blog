import {sendRequest} from './broker.js';

async function contactUs(email, subject, message) {
  console.log('contactUs');

  const payload = {
      action: 'mailContactUs',
      mail: {
        email: email,
        subject: subject,
        message: message
      }
  }

  return sendRequest(payload)
}

export { contactUs }