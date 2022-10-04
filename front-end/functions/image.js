import {sendRequestFormdata} from './broker.js';

async function uploadImage(token, data) {
  console.log('uploadImage');

  let formData = new FormData();
  formData.append("command", "uploadImage");
  formData.append("token", token);
  formData.append("file", data);

  return sendRequestFormdata(formData)
}

export { uploadImage }
