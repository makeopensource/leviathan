let urlParams = new URLSearchParams(window.location.search);
const jobId = urlParams.get("jobid");
console.log(jobId);

if (!jobId) {
    alert("Job ID is required");
}

const socket = new WebSocket(`ws://${document.location.hostname}:${document.location.port}/ws?jobid=${jobId}`);

// Define what happens when you receive a message
socket.onmessage = (event) => {
    const response = JSON.parse(event.data);
    console.log(response)

    let jobStatus = JSON.stringify(response.jobStatus)
    let logs = JSON.stringify(response.logs)

    console.log(logs)
    console.log(jobStatus)

    document.getElementById('job').innerHTML = `<p>Job Info: ${jobStatus}</p>`;
    document.getElementById('code').innerHTML = `<p>${logs}</p>`;
};

// Handle any errors that occur
socket.onerror = (error) => {
    console.error('WebSocket Error: ', error);
};

// Handle connection closure
socket.onclose = () => {
    console.log('WebSocket connection closed');
};
