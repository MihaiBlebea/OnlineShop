<template>
    <div>
        <p>{{ stream.length }} events - 
            <span>
                Show: 
                <span class="clickable" v-bind:class="{ 'is-blue': show === null }" v-on:click="setShowJust(null)">all</span> | 
                <span class="clickable" v-bind:class="{ 'is-blue': show === 25 }" v-on:click="setShowJust(25)">last 25</span> | 
                <span class="clickable" v-bind:class="{ 'is-blue': show === 50 }" v-on:click="setShowJust(50)">last 50</span>
            </span> - 
            <span>
                Order:
                <span class="clickable" v-bind:class="{ 'is-blue': orderDesc }" v-on:click="toggleOrder()">desc</span> |
                <span class="clickable" v-bind:class="{ 'is-blue': !orderDesc }" v-on:click="toggleOrder()">asc</span>
            </span>
                
        </p>
        <ul class="list-group">
            <li class="list-group-item" v-for="(event, index) in showJust" :key="index">
                <Event :event="event" />
            </li>
        </ul>
    </div>
</template>

<script>
import Event from './Event.vue'
import moment from 'moment'

export default {
    components: {
        Event
    },
    props: {
        stream: {
            type: Array,
            required: false,
            default: []
        }
    },
    data: function()
    {
        return {
            show: null,
            orderDesc: true
        }
    },
    computed: {
        orderedStream: function()
        {
            if(this.stream.length === 0)
            {
                return []
            }
            return this.stream.sort((a, b)=> {
                let aUtc = moment(a.timestamp).utc()
                let bUtc = moment(b.timestamp).utc()
                if(this.orderDesc)
                {
                    return aUtc - bUtc
                } else {
                    return bUtc - aUtc
                }
            })
        },
        showJust: function()
        {
            if(this.show === null)
            {
                return this.orderedStream
            }

            return this.orderedStream.slice(Math.max(this.orderedStream.length - this.show, 0))
        }
    },
    methods:
    {
        setShowJust: function(value)
        {
            this.show = value
        },
        toggleOrder: function()
        {
            this.orderDesc = !this.orderDesc
        }
    }
}
</script>

<style scoped>
.is-blue {
    color: blue;
}
.is-red {
    color: red;
}
.clickable {
    cursor: pointer;
}
</style>