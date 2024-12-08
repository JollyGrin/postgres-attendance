"use client";
import { useQuery } from "@tanstack/react-query";
import { AttendRecord, getDayDuration } from "./services/apiClient";
import { useDclPlayer } from "./services/dcl/useDclPlayer";

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

function getDuration(start: string, end: string) {
  const A = new Date(start).getTime();
  const B = new Date(end).getTime();
  return {
    durationMs: (B - A) / 1000,
    durationHr: (B - A) / 1000 / (60 * 60),
    first: A,
    last: B,
  };
}

const RecordList = () => {
  const searchDate = "2024-12-03";
  const { data } = useQuery({
    queryKey: ["duration", searchDate],
    queryFn: () => getDayDuration(searchDate),
  });

  const [firstStart] =
    data?.sort((a, b) => {
      const A = new Date(a.enter_time).getTime();
      const B = new Date(b.enter_time).getTime();
      return A - B;
    }) ?? [];

  const [lastEnd] =
    data?.sort((a, b) => {
      const A = new Date(a.exit_time).getTime();
      const B = new Date(b.exit_time).getTime();
      return B - A;
    }) ?? [];

  const { durationMs, first, last } = getDuration(
    firstStart?.enter_time,
    lastEnd?.exit_time,
  );

  return (
    <>
      {data
        // ?.splice(4, 1)
        ?.map((record) => (
          <Record
            key={record.address + record.enter_time}
            {...record}
            first={first}
            last={last}
            totalDuration={durationMs}
          />
        ))}
    </>
  );
};

const Record = (
  record: AttendRecord & {
    first?: number;
    last?: number;
    totalDuration?: number;
  },
) => {
  if (!record.first || !record.last || !record?.totalDuration) return null;
  const { data } = useDclPlayer(record.address);
  const [avatar] = data?.avatars ?? [];
  const pfp = avatar?.avatar.snapshots.face256;

  const percentageWidth = record.duration / record.totalDuration;
  const startDifference = new Date(record.enter_time).getTime() - record.first;
  const percentageLeft = startDifference / 1000 / record.totalDuration;
  console.table({
    total: record.totalDuration,
    duration: record.duration,
    percentageWidth,
    startTime: new Date(record.enter_time).getTime(),
    firstTime: record.first,
    startDifference: startDifference / 1000,
    percentageLeft,
  });

  return (
    <div className="grid grid-cols-[3rem_5rem_1fr] items-center">
      <img
        src={avatar?.avatar.snapshots.face256}
        key="pfp"
        style={{
          background: "black",
          minWidth: "2rem",
          minHeight: "2rem",
          borderRadius: "100%",
          margin: "0.25rem",
        }}
      />

      <p>{record.address.substring(0, 6)}...</p>
      <div className="bg-green-50 min-h-5 w-full relative">
        <div
          className="absolute bg-green-400 h-full"
          style={{
            width: `${percentageWidth * 100}%`,
            left: `${percentageLeft * 100}%`,
          }}
        ></div>
        {/* add bar that starts and ends at time */}
      </div>
    </div>
  );
};
