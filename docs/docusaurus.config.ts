import {themes as prismThemes} from 'prism-react-renderer';


// With JSDoc @type annotations, IDEs can provide config autocompletion
/** @type {import('@docusaurus/types').DocusaurusConfig} */
(module.exports = {
    title: 'Leviathan',
    tagline: 'Container orchestrator and job runner',
    url: 'https://makeopensource.github.io',
    baseUrl: '/leviathan/',
    onBrokenLinks: 'throw',
    onBrokenMarkdownLinks: 'warn',
    favicon: 'img/favicon.ico',
    organizationName: 'makeopensource',
    projectName: 'leviathan',

    presets: [
        [
            '@docusaurus/preset-classic',
            /** @type {import('@docusaurus/preset-classic').Options} */
            ({
                docs: {
                    sidebarPath: require.resolve('./sidebars.js'),
                    editUrl: 'https://github.com/makeopensource/leviathan/edit/main/',
                },
                blog: {
                    showReadingTime: true,
                    editUrl:
                        'https://github.com/makeopensource/leviathan/edit/main/blog/',
                },
                theme: {
                    customCss: require.resolve('./src/css/custom.css'),
                },
            }),
        ],
    ],

    themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
        ({
            navbar: {
                title: 'Leviathan',
                logo: {
                    alt: 'My Site Logo',
                    src: 'img/logo.svg',
                },
                items: [
                    {
                        type: 'doc',
                        docId: 'users/intro',
                        position: 'left',
                        label: 'User Guide',
                    },
                    {
                        type: 'doc',
                        docId: 'developers/intro',
                        position: 'left',
                        label: 'Developers Guide',
                    },
                    // remove this comment to re add blog
                    // {to: '/blog', label: 'Blog', position: 'left'},
                    {
                        href: 'https://github.com/makeopensource/leviathan',
                        label: 'GitHub',
                        position: 'right',
                    },
                ],
            },
            footer: {
                style: 'dark',
                links: [
                    {
                        title: 'Docs',
                        items: [
                            {
                                label: 'Tutorial',
                                to: '/docs/intro',
                            },
                        ],
                    },
                    {
                        title: 'Community',
                        items: [
                            {
                                label: 'Discord',
                                href: 'https://discord.gg/ChrT2DfcDT',
                            },
                        ],
                    },
                    {
                        title: 'More',
                        items: [
                            {
                                label: 'Blog',
                                to: '/blog',
                            },
                            {
                                label: 'GitHub',
                                href: 'https://github.com/makeopensource/leviathan',
                            },
                        ],
                    },
                ],
                copyright: `Copyright © ${new Date().getFullYear()} Leviathan. Built with Docusaurus.`,
            },
            prism: {
                theme: prismThemes.github,
                darkTheme: prismThemes.dracula,
            },
        }),
});
