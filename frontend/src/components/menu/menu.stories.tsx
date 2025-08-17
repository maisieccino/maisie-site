import type { Meta, StoryObj, Parameters } from "@storybook/react";

import { Menu } from ".";
import { P } from "../typography";

const meta: Meta<typeof Menu> = {
  component: Menu,
};

export default meta;

type Story = StoryObj<typeof Menu>;

export const ContainerComponent: Story = {
  render: () => (
    <Menu>
      <P>Something</P>
    </Menu>
  ),
};

