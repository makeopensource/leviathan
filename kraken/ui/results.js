let urlParams = new URLSearchParams(window.location.search);
const jobId = urlParams.get("jobid");
console.log(jobId);

if (!jobId) {
    alert("Job ID is required");
}

const socket = new WebSocket(`ws://${document.location.hostname}:${document.location.port}/ws?jobid=${jobId}`);
socket.onmessage = (event) => {
    const response = JSON.parse(event.data);
    console.log(response);

    // Display job status (prettified JSON)
    const jobStatusElement = document.getElementById('job');
    jobStatusElement.innerHTML = ""; // Clear previous content

    try {
        const jobStatus = response.jobStatus;
        const pre = document.createElement('pre'); // Use <pre> for formatting
        pre.textContent = JSON.stringify(jobStatus, null, 2); // Beautify JSON with indentation
        jobStatusElement.appendChild(pre);
    } catch (error) {
        jobStatusElement.innerHTML = "<p>Error displaying job status.</p>";
        console.error("Error parsing or displaying jobStatus:", error);
    }

    // Display logs (preserving newlines and handling special characters)
    const logsElement = document.getElementById('code');
    logsElement.innerHTML = ""; // Clear previous content

    try {
        const logs = response.logs;

        if (Array.isArray(logs)) {
            // Handle array of logs (common case)
            logs.forEach(logEntry => {
                const pre = document.createElement('pre');
                pre.textContent = logEntry; // Set text content to escape HTML
                logsElement.appendChild(pre);
            });
        } else if (typeof logs === 'string') {
            // Handle single log string
            const pre = document.createElement('pre');
            pre.textContent = logs;
            logsElement.appendChild(pre);
        } else {
            // Handle other log types if needed
            const pre = document.createElement('pre');
            pre.textContent = JSON.stringify(logs); // Or some other way to display
            logsElement.appendChild(pre);
        }


    } catch (error) {
        logsElement.innerHTML = "<p>Error displaying logs.</p>";
        console.error("Error displaying logs:", error);
    }
};

// // Define what happens when you receive a message
// socket.onmessage = (event) => {
//     const response = JSON.parse(event.data);
//     console.log(response)
//
//     let jobStatus = JSON.stringify(response.jobStatus)
//     let logs = JSON.stringify(response.logs)
//
//     console.log(logs)
//     console.log(jobStatus)
//
//     document.getElementById('job').innerHTML = `<p>Job Info: ${jobStatus}</p>`;
//     document.getElementById('code').innerHTML = `<p>${logs}</p>`;
// };

// Handle any errors that occur
socket.onerror = (error) => {
    console.error('WebSocket Error: ', error);
};

// Handle connection closure
socket.onclose = () => {
    console.log('WebSocket connection closed');
};
