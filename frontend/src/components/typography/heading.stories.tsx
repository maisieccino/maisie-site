import type { Meta, StoryObj } from "@storybook/react";

import { Heading } from ".";

const meta: Meta<typeof Heading> = {
  component: Heading,
};

export default meta;
type Story = StoryObj<typeof Heading>;

export const HeadingComponent: Story = {
  render: () => <Heading>Text</Heading>,
};
