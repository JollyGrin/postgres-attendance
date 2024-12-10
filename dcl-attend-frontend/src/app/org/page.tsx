import { IconChevDown } from "@/components/atoms/icons";
import Layout from "@/components/molecules/Layout";
import { OrgPanels } from "@/components/molecules/OrgPanels";

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
      <OrgPanels
        org={{ img: "logos/innkeeper.png", name: "innkeeper.eth" }}
        locations={[
          { location: "theinn", name: "hyperfy" },
          { location: "137,-2", name: "dcl" },
        ]}
      />
    </>
  );
};
