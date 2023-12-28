<script>

function userHasLiked(post, identifier) {
    return post.Likes.some(like => like.UserID == identifier);
}

export default {
    props: ["post", "identifier"],
    data: function () {
        return {
            liked: userHasLiked(this.post, this.identifier),
        };
    },
    methods: {
        toggleLike() {
            // console.log(this.identifier)
            // this.liked = !this.liked;
            // console.log("Like toggled");
            //if liked send delete request to photo/like/identifier else post, if successfull update liked
            if (this.liked) {
                this.$axios.delete("/photo/like/" + this.post.ID, { headers: { 'Authorization': this.identifier } })
                    .then((response) => {
                        console.log(response);
                        this.post.Likes = response.data;
                        this.liked = !this.liked;
                    })
                    .catch((error) => {
                        console.log(error);
                    })
            }
            else {
                this.$axios.post("/photo/like/" + this.post.ID, {}, { headers: { 'Authorization': this.identifier } })
                    .then((response) => {
                        console.log(response);
                        this.post.Likes = response.data;
                        this.liked = !this.liked;
                    })
                    .catch((error) => {
                        console.log(error);
                    })
            }
        }
    }
}
</script>

<template>
    <div v-if="post != null" class="post-container">
        <div class="username-container">
            <RouterLink :to="{path:'user/'+ post.UserID}" replace>
                <p>{{ post.UserName }}</p>
                <!-- <p>{{ this.test}}</p> -->
            </RouterLink>
        </div>
        <div class="image-container">
            <img :src="'http://localhost:3000/images/' + post.FileName" :alt="post.FileName">
        </div>
        <div class="likes-container">
            <p>
                <button class="like-btn" @click="toggleLike">
                    <svg class="feather" :class="this.liked ? 'liked': ''">
                        <use href="/feather-sprite-v4.29.0.svg#heart" />
                    </svg>
                </button>
                {{ post.Likes.length + ((post.Likes.length > 1) ? " likes" : " like") }}
            </p>
        </div>
        <div class="comments-container">
            <ul>
                <li class="comment" v-for="comment in post.Comments" :key="comment.ID">
                    {{ comment.Text }} - <small>{{ new Date(comment.CreateTime).toLocaleString() }}</small>
                </li>
            </ul>
            <div class="comment-input-container">
                <input type="text" placeholder="Comment" />
                <button>Post</button>
            </div>
        </div>
    </div>
</template>

<style>

.like-btn .feather.liked {
    fill: red;
    /* Color when liked */
}
</style>