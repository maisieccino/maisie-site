import { Meta, StoryObj } from "@storybook/react";

import { P } from "./p";

const meta: Meta<typeof P> = {
  component: P,
}

export default meta;
type Story = StoryObj<typeof P>;

export const PComponent: Story = {
  render: () => <>
    <P>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed facilisis metus quis viverra ornare. Suspendisse id ex non tortor venenatis egestas. Vestibulum iaculis, elit id volutpat tincidunt, ante ligula dignissim mauris, ut pretium ipsum neque ut nulla. Nunc sit amet dolor semper, hendrerit justo ut, convallis dui. Vestibulum id nisl eget nunc blandit vehicula ac vel risus. Nullam vel tortor eu ipsum euismod porta a ut ante. Cras blandit pulvinar volutpat. Sed id arcu eu nunc tempus luctus. Donec a maximus velit, in eleifend magna. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; In eleifend tincidunt sapien a tincidunt. Suspendisse potenti. Vivamus mollis risus vel aliquam commodo.</P>
    <P>Proin et tellus a libero luctus varius. Praesent nec velit vitae orci bibendum porttitor id eu lorem. Praesent vitae mi sed sapien pulvinar aliquam. Integer hendrerit blandit est ac finibus. Fusce sit amet mauris molestie, sagittis dui eu, imperdiet enim. Nunc semper pharetra ornare. Nullam nunc libero, euismod non risus ut, egestas pharetra nulla. Ut lorem diam, laoreet non sapien ut, eleifend vestibulum augue. Integer commodo nec nulla vitae pulvinar. Vestibulum blandit erat mauris, et dictum nunc malesuada aliquam. Integer eget nisi vehicula velit tincidunt dapibus eget quis mi.</P>
  </>
};
