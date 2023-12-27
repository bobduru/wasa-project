
  
<script>
export default {
    data() {
        return {
            name: ''
        };
    },
    methods: {
        async login() {
            try {
                // Send a post request using axios with access control allow origin header
                const response = await this.$axios.post('http://localhost:3000/login', { name: this.name }, { headers: { 'Access-Control-Allow-Origin': '*' } });
                // const response = await this.$axios.post('http://localhost:3000/login', { name: this.name });
                console.log(response.data);
                // setCookie('identifier', response.data.identifier);
                // this.$router.push('/home');
            } catch (error) {
                console.error(error);
            }
        }
    }
};
</script>
<template>
    <div class="login-container">

        <form @submit.prevent="login">
            <h2>Login</h2>
            <input type="text" v-model="name" placeholder="Enter your name" />
            <button type="submit">Submit</button>
        </form>
    </div>
</template>

  
<style>
.login-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
}

.login-container form {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.login-container input {
    margin: 10px;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
}

.login-container button {
    padding: 10px 20px;
    border: none;
    border-radius: 5px;
    background-color: #007bff;
    color: white;
    cursor: pointer;
}

.login-container button:hover {
    background-color: #0056b3;
}
</style>
  