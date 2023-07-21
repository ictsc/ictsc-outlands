import "@testing-library/jest-dom";
import { render, screen } from "@testing-library/react";
import { Mock, vi } from "vitest";

import useUserGroups from "@/hooks/userGroups";
import Users from "@/pages/users";
import { testUser } from "@/types/User";
import { testUserGroup } from "@/types/UserGroup";

vi.mock("next/error", () => ({
  __esModule: true,
  default: () => <div data-testid="error" />,
}));
vi.mock("@/hooks/userGroups");
vi.mock("@/components/Navbar", () => ({
  __esModule: true,
  default: () => <div data-testid="navbar" />,
}));
vi.mock("@/components/LoadingPage", () => ({
  __esModule: true,
  default: () => <div data-testid="loading" />,
}));

beforeEach(() => {
  // toHaveBeenCalledTimes がテストごとにリセットされるようにする
  vi.clearAllMocks();
});

describe("Users", () => {
  test("画面が表示されることを確認する", async () => {
    // setup
    (useUserGroups as Mock).mockReturnValue({
      userGroups: [{ ...testUserGroup, members: [testUser] }],
      isLoading: false,
    });

    // when
    render(<Users />);

    // then
    expect(screen.getByTestId("navbar")).toBeInTheDocument();
    const cells = screen.getAllByRole("cell");
    expect(cells[0]).toHaveTextContent(testUser.display_name);
    expect(cells[1]).toHaveTextContent(testUserGroup.name);
    expect(cells[2]).toHaveTextContent(testUser.profile?.self_introduction!);
    const links = screen.getAllByRole("link");
    expect(links[0]).toHaveAttribute(
      "href",
      `https://github.com/${testUser.profile?.github_id}`
    );
    expect(links[1]).toHaveAttribute(
      "href",
      `https://twitter.com/${testUser.profile?.twitter_id}`
    );
    expect(links[2]).toHaveAttribute(
      "href",
      `https://www.facebook.com/${testUser.profile?.facebook_id}`
    );

    // verify
    expect(useUserGroups).toHaveBeenCalledTimes(1);
  });

  test("通知が取得中の場合、ローディング画面が表示されることを確認する", async () => {
    // setup
    (useUserGroups as Mock).mockReturnValue({
      userGroups: [{ ...testUserGroup, members: [testUser] }],
      isLoading: true,
    });

    // when
    render(<Users />);

    // then
    expect(screen.getByTestId("navbar")).toBeInTheDocument();
    expect(screen.getByTestId("loading")).toBeInTheDocument();

    // verify
    expect(useUserGroups).toHaveBeenCalledTimes(1);
  });

  test("UserGroup が取得できなかった場合、エラーメッセージが表示されることを確認する", async () => {
    // setup
    (useUserGroups as Mock).mockReturnValue({
      userGroups: null,
      isLoading: false,
    });

    // when
    render(<Users />);

    // then
    expect(screen.getByTestId("error")).toBeInTheDocument();

    // verify
    expect(useUserGroups).toHaveBeenCalledTimes(1);
  });

  test("UserGroup が空の場合、テーブルに何も表示されないことを確認する", async () => {
    // setup
    (useUserGroups as Mock).mockReturnValue({
      userGroups: [],
      isLoading: false,
    });

    // when
    render(<Users />);

    // then
    expect(screen.getByTestId("navbar")).toBeInTheDocument();
    expect(screen.queryByRole("cell")).toBeNull();

    // verify
    expect(useUserGroups).toHaveBeenCalledTimes(1);
  });

  test("Member が空の場合、テーブルに何も表示されないことを確認する", async () => {
    // setup
    (useUserGroups as Mock).mockReturnValue({
      userGroups: [{ ...testUserGroup, members: [] }],
      isLoading: false,
    });

    // when
    render(<Users />);

    // then
    expect(screen.getByTestId("navbar")).toBeInTheDocument();
    expect(screen.queryByRole("cell")).toBeNull();
  });

  test("Github ID が未入力の場合、Github リンクが表示されないことを確認する", async () => {
    // setup
    const user = {
      ...testUser,
      profile: { ...testUser.profile, github_id: "" },
    };
    (useUserGroups as Mock).mockReturnValue({
      userGroups: [{ ...testUserGroup, members: [user] }],
      isLoading: false,
    });

    // when
    render(<Users />);

    // then
    expect(screen.getByTestId("navbar")).toBeInTheDocument();
    const links = screen.getAllByRole("link");
    expect(links[0]).toHaveAttribute(
      "href",
      `https://twitter.com/${testUser.profile?.twitter_id}`
    );
    expect(links[1]).toHaveAttribute(
      "href",
      `https://www.facebook.com/${testUser.profile?.facebook_id}`
    );

    // verify
    expect(useUserGroups).toHaveBeenCalledTimes(1);
  });

  test("Twitter ID が未入力の場合、Twitter リンクが表示されないことを確認する", async () => {
    // setup
    const user = {
      ...testUser,
      profile: { ...testUser.profile, twitter_id: "" },
    };
    (useUserGroups as Mock).mockReturnValue({
      userGroups: [{ ...testUserGroup, members: [user] }],
      isLoading: false,
    });

    // when
    render(<Users />);

    // then
    expect(screen.getByTestId("navbar")).toBeInTheDocument();
    const links = screen.getAllByRole("link");
    expect(links[0]).toHaveAttribute(
      "href",
      `https://github.com/${testUser.profile?.github_id}`
    );

    // verify
    expect(useUserGroups).toHaveBeenCalledTimes(1);
  });

  test("Facebook ID が未入力の場合、Facebook リンクが表示されないことを確認する", async () => {
    // setup
    const user = {
      ...testUser,
      profile: { ...testUser.profile, facebook_id: "" },
    };
    (useUserGroups as Mock).mockReturnValue({
      userGroups: [{ ...testUserGroup, members: [user] }],
      isLoading: false,
    });

    // when
    render(<Users />);

    // then
    expect(screen.getByTestId("navbar")).toBeInTheDocument();
    const links = screen.getAllByRole("link");
    expect(links[0]).toHaveAttribute(
      "href",
      `https://github.com/${testUser.profile?.github_id}`
    );
    expect(links[1]).toHaveAttribute(
      "href",
      `https://twitter.com/${testUser.profile?.twitter_id}`
    );

    // verify
    expect(useUserGroups).toHaveBeenCalledTimes(1);
  });

  test("プロフィールが存在しない場合、リンクが表示されないことを確認する", async () => {
    // setup
    const user = { ...testUser, profile: null };
    (useUserGroups as Mock).mockReturnValue({
      userGroups: [{ ...testUserGroup, members: [user] }],
      isLoading: false,
    });

    // when
    render(<Users />);

    // then
    expect(screen.getByTestId("navbar")).toBeInTheDocument();
    expect(screen.queryByRole("link")).not.toBeInTheDocument();
  });
});