"use client";
import { useQuery } from "@tanstack/react-query";
import { AttendRecord, getDayDuration } from "@/services/apiClient";
import { useDclPlayer } from "@/services/dcl/useDclPlayer";

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

export const RecordList = () => {
  const searchDate = "2024-12-08";
  const { data } = useQuery({
    queryKey: ["duration", searchDate],
    queryFn: () => getDayDuration(searchDate),
  });

  const attendees = data
    ?.filter(
      (record, index, self) =>
        self.findIndex((rec) => rec.address === record.address) === index,
    )
    ?.map((record) => ({
      address: record.address,
      durations: data.filter((rec) => rec.address === record.address),
    }));

  console.log({ attendees });

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
      <div className="grid grid-cols-[8rem_1fr] items-center">
        <div />
        <div className="flex justify-between">
          <p>{firstStart?.enter_time.split("T")?.[1]?.split(".")?.[0]}</p>
          <p>{lastEnd?.exit_time.split("T")?.[1]?.split(".")?.[0]}</p>
        </div>
      </div>
      {attendees
        // ?.splice(4, 1)
        ?.map((record) => (
          <Record
            key={record.address}
            address={record.address}
            durations={record.durations}
            first={first}
            last={last}
            totalDuration={durationMs}
          />
        ))}
    </>
  );
};

const Record = (record: {
  address: string;
  durations: AttendRecord[];
  first?: number;
  last?: number;
  totalDuration?: number;
}) => {
  if (!record.first || !record.last || !record?.totalDuration) return null;
  const { data } = useDclPlayer(record.address);
  const [avatar] = data?.avatars ?? [];

  function findLeft(enterTime: string) {
    if (!enterTime || !record.first || !record.totalDuration) return 0;
    const startDifference = new Date(enterTime).getTime() - record.first;
    const percentageLeft = startDifference / 1000 / record.totalDuration;
    return percentageLeft;
  }
  // console.table({
  //   total: record.totalDuration,
  //   duration: record.duration,
  //   percentageWidth,
  //   startTime: new Date(record.enter_time).getTime(),
  //   firstTime: record.first,
  //   startDifference: startDifference / 1000,
  //   percentageLeft,
  // });

  return (
    <div className="grid grid-cols-[3rem_5rem_1fr] items-center">
      <img
        src={avatar?.avatar.snapshots.face256}
        key="pfp"
        style={{
          background: "black",
          width: "2.5rem",
          height: "2.5rem",
          borderRadius: "100%",
          margin: "0.25rem",
        }}
      />

      <p>{record.address.substring(0, 6)}...</p>
      <div className="bg-green-50 min-h-5 w-full relative rounded-xl">
        {record?.durations?.map((dur) => (
          <ChartLine
            key={dur.address + dur.enter_time}
            width={dur.duration / (record?.totalDuration ?? 1)}
            left={findLeft(dur?.enter_time)}
          />
        ))}
      </div>
    </div>
  );
};

const ChartLine = (props: { width: number; left: number }) => {
  return (
    <div
      className="absolute bg-green-400 h-full rounded-xl"
      style={{
        width: `${props.width * 100}%`,
        left: `${props.left * 100}%`,
      }}
    ></div>
  );
};
