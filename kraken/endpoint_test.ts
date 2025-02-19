// import fetch from 'node-fetch';
// import FormData from 'form-data';
// import fs from 'fs/promises';
// import {describe, it} from "node:test"; // For reading files
//
// const API_ENDPOINT = 'http://localhost:3000/your-job-endpoint'; // Replace with your actual endpoint
//
// describe('Job Endpoint (Integration Test)', () => {
//     it('should create a new job successfully', async () => {
//         const formData = new FormData();
//         formData.append('imageTag', 'test-image');
//         formData.append('timeoutInSeconds', '60');
//
//         // Attach files.  Adjust paths as needed.
//         formData.append('grader', await fs.readFile('./__tests__/fixtures/grader.txt'), 'grader.txt'); // Use real files for integration tests
//         formData.append('makefile', await fs.readFile('./__tests__/fixtures/makefile'), 'makefile');
//         formData.append('student', await fs.readFile('./__tests__/fixtures/student.zip'), 'student.zip');
//         formData.append('dockerfile', await fs.readFile('./__tests__/fixtures/Dockerfile'), 'Dockerfile');
//
//
//         const response = await fetch(API_ENDPOINT, {
//             method: 'POST',
//             body: formData,
//         });
//
//         expect(response.status).toBe(200); // Or expect(response.status).toBe(302) if you're redirecting
//         // If you're redirecting and want to follow the redirect
//         // const redirectedResponse = await fetch(response.headers.get('location') as string);
//         // expect(redirectedResponse.status).toBe(200); // Check the status of the redirected page
//         // ... other assertions on the redirected page if needed
//     });
//     // ... other integration tests as needed (e.g., handling missing files, server errors)
// });