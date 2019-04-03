#pragma checksum "/Users/Lily/Projects/BudgetPlanning/BudgetPlanning/Views/Home/Dashboard.cshtml" "{ff1816ec-aa5e-4d10-87f7-6f4963833460}" "16687d8cb5c57c16e712e28fa8c6c84796750561"
// <auto-generated/>
#pragma warning disable 1591
[assembly: global::Microsoft.AspNetCore.Razor.Hosting.RazorCompiledItemAttribute(typeof(AspNetCore.Views_Home_Dashboard), @"mvc.1.0.view", @"/Views/Home/Dashboard.cshtml")]
[assembly:global::Microsoft.AspNetCore.Mvc.Razor.Compilation.RazorViewAttribute(@"/Views/Home/Dashboard.cshtml", typeof(AspNetCore.Views_Home_Dashboard))]
namespace AspNetCore
{
    #line hidden
    using System;
    using System.Collections.Generic;
    using System.Linq;
    using System.Threading.Tasks;
    using Microsoft.AspNetCore.Mvc;
    using Microsoft.AspNetCore.Mvc.Rendering;
    using Microsoft.AspNetCore.Mvc.ViewFeatures;
#line 1 "/Users/Lily/Projects/BudgetPlanning/BudgetPlanning/Views/_ViewImports.cshtml"
using BudgetPlanning;

#line default
#line hidden
#line 2 "/Users/Lily/Projects/BudgetPlanning/BudgetPlanning/Views/_ViewImports.cshtml"
using BudgetPlanning.Models;

#line default
#line hidden
    [global::Microsoft.AspNetCore.Razor.Hosting.RazorSourceChecksumAttribute(@"SHA1", @"16687d8cb5c57c16e712e28fa8c6c84796750561", @"/Views/Home/Dashboard.cshtml")]
    [global::Microsoft.AspNetCore.Razor.Hosting.RazorSourceChecksumAttribute(@"SHA1", @"4143aed6b9ff032f9234ea78c1fa3fd5d1a8919f", @"/Views/_ViewImports.cshtml")]
    public class Views_Home_Dashboard : global::Microsoft.AspNetCore.Mvc.Razor.RazorPage<dynamic>
    {
        private static readonly global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute __tagHelperAttribute_0 = new global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute("style", new global::Microsoft.AspNetCore.Html.HtmlString("background-color:#f2f2f2;"), global::Microsoft.AspNetCore.Razor.TagHelpers.HtmlAttributeValueStyle.DoubleQuotes);
        #line hidden
        #pragma warning disable 0169
        private string __tagHelperStringValueBuffer;
        #pragma warning restore 0169
        private global::Microsoft.AspNetCore.Razor.Runtime.TagHelpers.TagHelperExecutionContext __tagHelperExecutionContext;
        private global::Microsoft.AspNetCore.Razor.Runtime.TagHelpers.TagHelperRunner __tagHelperRunner = new global::Microsoft.AspNetCore.Razor.Runtime.TagHelpers.TagHelperRunner();
        private global::Microsoft.AspNetCore.Razor.Runtime.TagHelpers.TagHelperScopeManager __backed__tagHelperScopeManager = null;
        private global::Microsoft.AspNetCore.Razor.Runtime.TagHelpers.TagHelperScopeManager __tagHelperScopeManager
        {
            get
            {
                if (__backed__tagHelperScopeManager == null)
                {
                    __backed__tagHelperScopeManager = new global::Microsoft.AspNetCore.Razor.Runtime.TagHelpers.TagHelperScopeManager(StartTagHelperWritingScope, EndTagHelperWritingScope);
                }
                return __backed__tagHelperScopeManager;
            }
        }
        private global::Microsoft.AspNetCore.Mvc.Razor.TagHelpers.BodyTagHelper __Microsoft_AspNetCore_Mvc_Razor_TagHelpers_BodyTagHelper;
        #pragma warning disable 1998
        public async override global::System.Threading.Tasks.Task ExecuteAsync()
        {
#line 1 "/Users/Lily/Projects/BudgetPlanning/BudgetPlanning/Views/Home/Dashboard.cshtml"
  
    ViewData["Title"] = "My Dashboard";

#line default
#line hidden
            BeginContext(48, 2443, true);
            WriteLiteral(@"
<!DOCTYPE html>
<title>Budget Planning</title>
<link rel=""stylesheet"" href=""https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css"">
<script src=""https://cdn.anychart.com/js/8.0.1/anychart-core.min.js""></script>
<script src=""https://cdn.anychart.com/js/8.0.1/anychart-pie.min.js""></script>

<style>
    p, h1, h2, h3, h4, h5 {
        font-family: sans-serif;
    }
    .square {
        width: 150px;
        height: 150px;
        background-color: #2ecc71;
        color: white;
        display: inline-block;
        border: none;
        vertical-align: top;
        float: left;
        margin-right: 10px;
    }

        .square:hover {
            background-color: #68d496;
            animation-timing-function: ease-in;
            animation-duration: 50ms;
        }

    .squareB {
        width: 150px;
        height: 150px;
        background-color: #0891eb;
        color: white;
        display: inline-block;
        border: none;
        padding-right:");
            WriteLiteral(@" 10px;
        vertical-align: top;
        float: left;
        margin-right: 10px;
    }

        .squareB:hover {
            background-color: #37a7f0;
            animation-timing-function: ease-in;
            animation-duration: 50ms;
        }

    .squareC {
        width: 150px;
        height: 150px;
        background-color: #e74c3c;
        color: white;
        display: inline-block;
        border: none;
        padding-right: 10px;
        vertical-align: top;
        float: left;
        margin-right: 10px;
    }

        .squareC:hover {
            background-color: #ea7063;
            animation-timing-function: ease-in;
            animation-duration: 50ms;
        }

    .squareD {
        width: 150px;
        height: 150px;
        background-color: #8c37f9;
        color: white;
        display: inline-block;
        border: none;
        padding-right: 10px;
        vertical-align: top;
        float: left;
        margin-right: 10px;
    }
");
            WriteLiteral(@"
        .squareD:hover {
            background-color: #9e52ff;
            animation-timing-function: ease-in;
            animation-duration: 50ms;
        }

    .cont {
        height: 310px;
        width: 500px;
        background-color: white;
        display: inline-block;
        vertical-align: top;
        float: left;
        margin-right: 10px;
    }
</style>

");
            EndContext();
            BeginContext(2491, 2832, false);
            __tagHelperExecutionContext = __tagHelperScopeManager.Begin("body", global::Microsoft.AspNetCore.Razor.TagHelpers.TagMode.StartTagAndEndTag, "e9eb71d39c744980885240f9f4c49ef3", async() => {
                BeginContext(2531, 2785, true);
                WriteLiteral(@"
    <h3></h3>
    <div>
        <div class=""row"">
            <div class=""cont"" style=""overflow:hidden;"">
                <h3 style=""padding-left:20px;"">Spending Categories</h3>
                <div id=""container"" style=""width: 90%; height: 90%;margin-top:-5px;""></div>
                <script>
                    anychart.onDocumentReady(function() {

                    var data = [
                        {x: ""Food"", value: 45.50},
                        {x: ""Transportation"", value: 24.49},
                        {x: ""Shopping"", value: 100.44},
                        {x: ""Miscellaneous"", value: 2.79},
                    ];

                    var chart = anychart.pie();

                    chart.legend().position(""right"");
                    chart.legend().itemsLayout(""vertical"");
                    chart.data(data);
                    chart.container('container');
                    chart.draw();

                    });
                </script>
            </div>
  ");
                WriteLiteral(@"          <button class=""square"">Log Spendings</button>
            <button class=""squareB"">Fixed Expenses</button>
            <button class=""squareD"">Timeline</button>
            <button class=""squareC"">Settings</button>
            <div class=""cont"" style=""height:150px;width:310px;margin-top:10px;"">
                <div style=""text-align:center;padding-left:15px;padding-right:15px;"">
                    <h3>You have spent $156.00 in the past month.</h3>
                </div>
            </div>
            <div class=""cont"" style=""height:150px;width:310px;margin-top:10px;"">
                <div style=""padding-left:15px;"">
                    <h4>Other Links</h4>
                    <div style=""color:cadetblue"">
                        <p>Users</p>
                        <p>Abp Users</p>
                        <p>Report Feedback</p>
                    </div>
                </div>
            </div>
        </div>
        <div class=""row"">
            <div class=""cont"" style=""width:");
                WriteLiteral(@"565px;margin-top:10px;"">
                <div style=""padding-left:20px;"">
                    <h3>Transactions</h3>
                    <br />
                    <h5>Preview Transactions from Table here...</h5>
                </div>

            </div>
            <div class=""cont"" style=""width:565px;margin-top:10px;"">
                <div style=""padding-left:20px;"">
                    <h3>Top Categories of the Month</h3>
                    <br />
                    <h5>1 ----- Transportation</h5>
                    <h5>2 ----- Shopping</h5>
                    <h5>3 ----- Food</h5>
                    <h5>4 ----- Miscellaneous</h5>
                </div>
            </div>

        </div>
    </div>
");
                EndContext();
            }
            );
            __Microsoft_AspNetCore_Mvc_Razor_TagHelpers_BodyTagHelper = CreateTagHelper<global::Microsoft.AspNetCore.Mvc.Razor.TagHelpers.BodyTagHelper>();
            __tagHelperExecutionContext.Add(__Microsoft_AspNetCore_Mvc_Razor_TagHelpers_BodyTagHelper);
            __tagHelperExecutionContext.AddHtmlAttribute(__tagHelperAttribute_0);
            await __tagHelperRunner.RunAsync(__tagHelperExecutionContext);
            if (!__tagHelperExecutionContext.Output.IsContentModified)
            {
                await __tagHelperExecutionContext.SetOutputContentAsync();
            }
            Write(__tagHelperExecutionContext.Output);
            __tagHelperExecutionContext = __tagHelperScopeManager.End();
            EndContext();
            BeginContext(5323, 6, true);
            WriteLiteral("\r\n\r\n\r\n");
            EndContext();
        }
        #pragma warning restore 1998
        [global::Microsoft.AspNetCore.Mvc.Razor.Internal.RazorInjectAttribute]
        public global::Microsoft.AspNetCore.Mvc.ViewFeatures.IModelExpressionProvider ModelExpressionProvider { get; private set; }
        [global::Microsoft.AspNetCore.Mvc.Razor.Internal.RazorInjectAttribute]
        public global::Microsoft.AspNetCore.Mvc.IUrlHelper Url { get; private set; }
        [global::Microsoft.AspNetCore.Mvc.Razor.Internal.RazorInjectAttribute]
        public global::Microsoft.AspNetCore.Mvc.IViewComponentHelper Component { get; private set; }
        [global::Microsoft.AspNetCore.Mvc.Razor.Internal.RazorInjectAttribute]
        public global::Microsoft.AspNetCore.Mvc.Rendering.IJsonHelper Json { get; private set; }
        [global::Microsoft.AspNetCore.Mvc.Razor.Internal.RazorInjectAttribute]
        public global::Microsoft.AspNetCore.Mvc.Rendering.IHtmlHelper<dynamic> Html { get; private set; }
    }
}
#pragma warning restore 1591