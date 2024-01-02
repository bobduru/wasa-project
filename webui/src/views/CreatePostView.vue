<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { getCookie } from '../utils/cookieUtils';
const route = useRouter();


</script>
<script>
export default {
    data() {
        return {
            image:null
        };
    },
    methods: {
        handleFileChange(event) {
            this.image = event.target.files[0];
        },
        async submitPost() {
            if (!this.image) {
                alert("Please select a file.");
                return;
            }

            const identifier = getCookie('identifier');
            const formData = new FormData();
            formData.append('image', this.image);

            try {
                await this.$axios.post('http://localhost:3000/photo', formData, {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                        'Authorization': identifier
                    }
                });
                alert("File Uploaded successfully");
                this.router.push('/');
            } catch (error) {
                console.error(error);
            }
        }
    }
}
</script>

<template>
    <div class="create-post-container">
        <form @submit.prevent="submitPost">
            <h2>Create Post</h2>
            <input type="file" @change="handleFileChange" />
            <button type="submit">Submit</button>
        </form>
    </div>
</template>
  
  
  
<style>
.create-post-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    padding: 20px;
    background-color: #f4f4f4;
}

.create-post-container form {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 20px;
    background: white;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    width: 300px;
}

.create-post-container input[type="file"] {
    margin: 15px 0;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    width: 100%;
}

.create-post-container button {
    padding: 10px 20px;
    border: none;
    border-radius: 5px;
    background-color: #42b983;
    color: white;
    cursor: pointer;
    width: 100%;
    margin-top: 10px;
}

.create-post-container button:hover {
    background-color: #269261;
}
</style>
  