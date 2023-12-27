<script>
export default {
    props: ["post"],
    data: function () {
        return {
            liked: false,
        };
    },
    methods: {
        toggleLike() {
            this.liked = !this.liked;
            console.log("Like toggled");
        }
    }
}
</script>

<template>
    <div v-if="post != null" class="post-container">
        <div class="username-container">
            <p>{{ post.UserName }}</p>
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