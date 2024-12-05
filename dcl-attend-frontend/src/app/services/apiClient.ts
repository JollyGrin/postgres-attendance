import axios from "axios";

const baseUrl = "http://localhost:8080";

export async function getDayDuration(day: string) {
  const res = await axios.get(`${baseUrl}/api/attendance/duration?day=${day}`);
  return res.data;
}
