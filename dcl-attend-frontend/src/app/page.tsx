"use client";

import { IconChevDown } from "@/components/atoms/icons";
import { useState } from "react";

export default function Home() {
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
        className="border-border border-r-2 group-hover:[grid-template-columns:8rem_1fr]"
      ></div>
      <div className="flex flex-col">
        <div className="border-border border-b-2 min-h-16 px-4 flex items-center gap-3 font-mono">
          <span>org</span>
          <span className="text-border font-bold">/</span>
          <span>innkeeper</span>
          <IconChevDown color="var(--color-border)" />
        </div>
      </div>
    </div>
  );
}
