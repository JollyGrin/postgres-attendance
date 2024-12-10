import { IconChevDown } from "@/components/atoms/icons";
import Layout from "@/components/molecules/Layout";

export default function OrgPage() {
  return <Layout top={<Top />} body={<Body />} />;
}

const Top = () => (
  <>
    <span>org</span>
    <span className="text-border font-bold">/</span>
    <span>innkeeper</span>
    <IconChevDown color="var(--color-border)" />
  </>
);

const Body = () => {
  return (
    <>
      <div className="flex justify-between">
        <div className="bg-panel p-4 border-border border-2 rounded-lg flex items-center gap-4 text-xl">
          <img src="logos/innkeeper.png" height="40px" width="40px" />
          <p>innkeeper.eth</p>
        </div>

        <div className="flex gap-4 font-mono">
          <div className="bg-panel p-4 border-border border-2 rounded-lg flex items-center gap-4 text-xl">
            <img src="logos/hyperfy.png" height="40px" width="40px" />
            <div>
              <p className="text-xs">Hyperfy</p>
              <p className="text-md">theinn</p>
            </div>
          </div>

          <div className="bg-panel p-4 border-border border-2 rounded-lg flex items-center gap-4 text-xl">
            <img src="logos/dcl.png" height="40px" width="40px" />
            <div>
              <p className="text-xs">Decentraland</p>
              <p className="text-md">137,-2</p>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};
