import axios from 'axios';

 async function sendRequest(payload) {
    var result = null
  
    const headers = new Headers();
    headers.append("Content-Type", "application/json")
  
    const body = {
        method: 'POST',
        body: JSON.stringify(payload),
        headers: headers,
    }
  
    await fetch("http://localhost:8080/handle", body)
        .then((response) => response.json())
        .then((data) => {
            JSON.stringify(data, undefined, 4);
            if (data.error) {
                console.log('Error: ');
            } else {
                if(data.data){
                    result  = JSON.parse(data.data)
                } else {
                    result = true
                }
            }
        })
        .catch((error) => {
            console.log('Catch error: ' + error)
        })
  
    return result
  }

  async function sendRequestFormdata(formData) {
    var result = null

    await axios
        .post("http://localhost:8080/handleData", formData)
        .then((data) => {
            JSON.stringify(data, undefined, 4);
            if (data.error) {
                console.log('Error: ');
            } else {
                if(data.data.data){
                    result  = JSON.parse(data.data.data)
                } else {
                    result = true
                }
            }
        })
        .catch(error => {
            console.log(error);
        });
  
    return result
  }

  export { sendRequest, sendRequestFormdata }