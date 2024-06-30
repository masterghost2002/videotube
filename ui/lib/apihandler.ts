import axios from "axios";

const serverEndpoint = process.env.NEXT_PUBLIC_SERVER_ENDPOINT;
const VIDEO_TUBE = axios.create({
    baseURL:serverEndpoint,
})
export default VIDEO_TUBE