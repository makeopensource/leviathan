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

    const jobStatus = response.jobStatus;
    updateJobStatus(jobStatus);

    const stdout = response.logs;
    updateStdout(stdout);
};

function updateJobStatus(data) {
    try {
        // Update status badge
        const statusBadge = document.querySelector('.status-badge');
        statusBadge.textContent = data.status;

        // Update status message if it exists
        const statusMessage = document.querySelector('.status-message');
        if (data.statusMessage) {
            statusMessage.innerHTML = `<pre>${data.statusMessage}</pre>`;
            statusMessage.style.display = 'block';
        } else {
            statusMessage.style.display = 'none';
        }


        // Update all info values
        const infoMapping = {
            'Job ID': data.jobId,
            'Container ID': data.containerId,
        };

        // Update each field in the info grid
        Object.entries(infoMapping).forEach(([label, value]) => {
            // Find the label element
            const labelElements = Array.from(document.querySelectorAll('.info-label'));
            const labelEl = labelElements.find(el => el.textContent === label);

            if (labelEl) {
                // Get the corresponding value element (next sibling in grid)
                const valueEl = labelEl.nextElementSibling;
                if (valueEl) {
                    if (value && value.trim() !== '') {
                        valueEl.textContent = value;
                        valueEl.classList.remove('empty-value');
                    } else {
                        valueEl.textContent = `No ${label.toLowerCase()} specified`;
                        valueEl.classList.add('empty-value');
                    }
                }
            }
        });

        // Update status colors based on status
        updateStatusColors(data.status);

    } catch (error) {
        console.error('Error updating job status:', error);
        // Optionally show error state in UI
    }
}

function updateStatusColors(status) {
    const statusBadge = document.querySelector('.status-badge');
    const statusMessage = document.querySelector('.status-message');

    // Reset classes
    statusBadge.style.backgroundColor = '';
    statusBadge.style.color = '';
    statusMessage.style.backgroundColor = '';
    statusMessage.style.borderColor = '';
    statusMessage.style.color = '';

    // Apply colors based on status
    switch (status.toLowerCase()) {
        case 'failed':
            statusBadge.style.backgroundColor = '#431418';
            statusBadge.style.color = '#f87171';
            statusMessage.style.backgroundColor = '#2d1517';
            statusMessage.style.borderColor = '#431418';
            statusMessage.style.color = '#f87171';
            break;

        case 'complete':
            statusBadge.style.backgroundColor = '#064e3b';
            statusBadge.style.color = '#34d399';
            statusMessage.style.backgroundColor = '#064e3b';
            statusMessage.style.borderColor = '#065f46';
            statusMessage.style.color = '#34d399';
            break;

        case 'running':
            statusBadge.style.backgroundColor = '#1e40af';
            statusBadge.style.color = '#60a5fa';
            statusMessage.style.backgroundColor = '#1e3a8a';
            statusMessage.style.borderColor = '#1e40af';
            statusMessage.style.color = '#60a5fa';
            break;

        default:
            statusBadge.style.backgroundColor = '#374151';
            statusBadge.style.color = '#9ca3af';
            statusMessage.style.backgroundColor = '#1f2937';
            statusMessage.style.borderColor = '#374151';
            statusMessage.style.color = '#9ca3af';
    }
}

let autoScroll = true;

function updateStdout(stdout) {
    try {
        const container = document.getElementById('stdoutContainer');
        container.innerHTML = '';

        const lines = stdout.split('\n');
        for (let li of lines) {
            const lineElement = document.createElement('div');
            lineElement.className = 'stdout-line';
            lineElement.textContent = li;
            container.appendChild(lineElement);
        }

        // Auto-scroll to bottom if enabled
        if (autoScroll) {
            container.scrollTop = container.scrollHeight;
        }
    } catch (error) {
        console.error('Error updating stdout:', error);
    }
}

// Handle any errors that occur
socket.onerror = (error) => {
    console.error('WebSocket Error: ', error);
};

// Handle connection closure
socket.onclose = () => {
    console.log('WebSocket connection closed');
};
