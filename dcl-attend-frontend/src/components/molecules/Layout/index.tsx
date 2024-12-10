"use client";

import Link from "next/link";
import { ReactNode, useState } from "react";

export default function Layout(props: { top?: ReactNode; body?: ReactNode }) {
  const [isHovering, setIsHovering] = useState(false);
  const over = () => setIsHovering(true);
  const out = () => setIsHovering(false);

  return (
    <div
      className={`grid w-svh h-svh transition-all`}
      style={{
        gridTemplateColumns: `${isHovering ? 8 : 5}rem 1fr`,
      }}
    >
      <div
        onMouseOver={over}
        onMouseOut={out}
        className="border-border border-r-2 group-hover:[grid-template-columns:8rem_1fr] flex flex-col items-center"
      >
        <Link href="/">
          <div className="w-[40px] h-[40px] bg-border rounded-full mt-4" />
        </Link>
      </div>
      <div className="flex flex-col">
        <div className="border-border border-b-2 min-h-16 px-4 flex items-center gap-3 font-mono">
          {props.top}
        </div>
        <div className="p-4">{props.body}</div>
      </div>
    </div>
  );
}

/**

          <span>org</span>
          <span className="text-border font-bold">/</span>
          <span>innkeeper</span>
          <IconChevDown color="var(--color-border)" />
 
 * */
