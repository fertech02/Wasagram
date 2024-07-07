
<script>
import ErrorMsg from '@/components/ErrorMsg.vue'
const token = sessionStorage.getItem('authToken');
export default {
    components: {
        ErrorMsg
    },
    data() {
        return {
            newname: '',
            errore: false,
            error_msg: '',
        };
    },
    methods: {
        
        async setMyUserName() {
            if (this.newname === ''){
                this.errore = true;
                this.error_msg = 'Please enter a username';
            } else {
                try {

                    const response = await this.$axios.post('/users/' + this.$token + '/username/',
                        {
                            username: this.newname
                        },
                        {
                            headers: {
                                'Authorization': `Bearer ${token}`
                            }
                        }
                    );
                    console.log("Username changed successfully")
 
                } catch (e) {
                    if (e.response && e.response.status === 400) {
                        this.errore = true;
                        this.error_msg = 'Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.';
                    } else if (e.response && e.response.status === 500) {
                        this.errore = true;
                        this.error_msg = 'An internal error occurred. We will be notified. Please try again later.';
                    } else {
                        this.errore = true;
                        this.error_msg = e.toString();
                    }
                }   
            }
    }
    }
}
</script>

<template>
    <div class="container d-flex justify-content-center align-items-center vh-100">
        <form @submit.prevent="setMyUserName" class="border p-4 rounded">
            <h2 class="mb-4">Change your username</h2>
            <div class="mb-3">
                <label for="inputName" class="form-label">New Name</label>
                <input v-model="newname" type="text" class="form-control" id="inputName" required minlength="3" maxlength="16">
            </div>
            <button type="submit" class="btn btn-primary">Submit</button>
            <div class="alert alert-success" role="alert" v-if="changedSuccess" style="margin: 10px;">
                Name changed successfully!
            </div>
            <ErrorMsg :msg="error_msg" v-else-if="errore" style="margin: 10px;"/>
        </form>
    </div>
</template>