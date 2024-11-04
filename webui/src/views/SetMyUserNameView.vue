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
            changedSuccess: false,
            errore: false,
            error_msg: '',
        };
    },
    methods: {
        async submitForm() {
            try {
                const config = {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`,
                    },
                };
                const response = await this.$axios.put(`/users/${token}`, { username: this.newname }, config);
                console.log("Name changed");
                this.changedSuccess = true;
                this.errore = false;
            }
            catch (error) {
                console.error(error, "Error in changing name");
                const statusCode = error.response.status;
                switch (statusCode) {
                    case 401:
                        console.error('Access Unauthorized:', error.response.data);
                        this.error_msg = "You are not logged in"
                        break;
                    case 400:
                        console.error('Bad request:', error.response.data);
                        this.error_msg = "Name already in use"
                        break;
                    case 404:
                        console.error('Not found: ', error.response.data);
                        this.error_msg = "You are not logged in"
                    default:
                        console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
                        this.error_msg = "You should login first"
                }
                this.changedSuccess = false;
                this.errore = true;
            }

            this.newname = '';

        },
    },
};
</script>