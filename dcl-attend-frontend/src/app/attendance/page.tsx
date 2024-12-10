"use client";

import { RecordList } from "@/components/molecules/AttendanceDurations";

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

      <div className="bg-foreground text-background rounded mt-4 p-4 flex flex-col gap-4">
        <RecordList />
      </div>
    </div>
  );
}
