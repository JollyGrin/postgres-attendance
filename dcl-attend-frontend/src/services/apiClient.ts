import axios, { AxiosResponse } from "axios";

const baseUrl = "http://lorekeeper.xyz";

export type AttendRecord = {
  address: string;
  enter_time: string;
  exit_time: string;
  duration: number;
};

export async function getDayDuration(day: string) {
  const res = await axios.get<AxiosResponse<AttendRecord[]>>(
    `${baseUrl}/api/attendance/duration?day=${day}`,
  );
  return res.data.data;
}
