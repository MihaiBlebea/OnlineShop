<template>
    <div class="media">
        <img :src="realImgUrl(image)" class="align-self-start mr-3" alt="">
        <div class="media-body">
            <h5 class="mt-0">{{ title }} {{ quantityFormat }}</h5>
            <p>£{{ price }}</p>
            <p>{{ description }}</p>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        title: {
            required: true,
            type: String
        },
        price: {
            required: true,
            type: String
        },
        image: {
            required: true,
            type: String
        },
        quantity: {
            required: true,
            type: String
        },
        description: {
            required: true,
            type: String
        }
    },
    computed:
    {
        quantityFormat: function()
        {
            if(this.quantity > 1)
            {
                return `x${ this.quantity }`
            }
        }
    },
    methods:
    {
        realImgUrl: function(imgFile)
        {
            return `https://github.com/wedeploy-examples/supermarket-web-example/blob/master/ui/assets/images/${ imgFile }?raw=true`
        }
    }
}
</script>

<style scoped>
img {
    width: 64px;
    height: 64px;
}
</style>