<script>
import { getCookie } from '../utils/cookieUtils';
import { useRouter } from 'vue-router';

export default {
    props: {
        redirects: {
            type: Boolean,
            default: true
        }
    },
    setup() {
        const router = useRouter();
        return { router };

    },
    computed: {
        isAuthenticated() {
            return !!getCookie('identifier');
        }
    },
    created() {
        if (!this.isAuthenticated && this.redirects) {
            this.router.push('/login');
        }
    }
};
</script>

<template>
    <div>
        <slot v-if="isAuthenticated" />
        <div v-else>
            <span v-if="redirects">Redirecting to login...</span>
            <!-- <span v-else>Redirecting to login...</span> -->
        </div>
    </div>
</template>
