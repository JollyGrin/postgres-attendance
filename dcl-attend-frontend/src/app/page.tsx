"use client";
import { useQuery } from "@tanstack/react-query";
import { getDayDuration } from "./services/apiClient";

export default function Home() {
  return (
    <div className="h-screen max-w-[900px] mx-auto py-4">
      <div className="bg-foreground min-h-14 w-full rounded flex text-background items-center p-4 text-xl">
        <div>{"<"}</div>
        <div className="flex-grow justify-items-center">
          <p>Today</p>
        </div>
        <div>{">"}</div>
      </div>

      <RecordList />
    </div>
  );
}

const RecordList = () => {
  const { data } = useQuery({
    queryKey: ["duration"],
    queryFn: () => getDayDuration("2024-12-03"),
  });

  return (
    <div className="bg-foreground text-background rounded mt-4 p-4 flex flex-col gap-4">
      {data?.data?.map((record) => (
        <div key={record?.address + record?.enter_time}>
          <p>{record?.address}</p>
          <p>{record.enter_time}</p>
          <p>{Math.floor((record?.duration / 60) * 100) / 100} minutes</p>
        </div>
      ))}
    </div>
  );
};
